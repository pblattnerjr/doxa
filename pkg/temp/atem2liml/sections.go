package atem2mt

// TODO: Delete this file
import (
	"fmt"
	"os"
	"text/scanner"
)

// TPath provides the path to a template and the template name.
type TPath struct {
	Path         string
	TemplateName string
}

// Section provides a slice for which templates insert it and which declares it.
// Because a section name is prefixed by a template name, it is possible for the
// same name to be declared by multiple templates.
type Section struct {
	Name          string
	WhereInserted []TPath
	WhereDeclared []TPath
}

// AddInsert adds a template path to the list of where this section is inserted.
func (s *Section) AddInsert(path TPath) {
	s.WhereInserted = append(s.WhereInserted, path)
}

// AddWhereDeclared adds a template path to the list of where this section is declared.
// Multiple templates can have sections of the same name.
func (s *Section) AddWhereDeclared(path TPath) {
	s.WhereDeclared = append(s.WhereDeclared, path)
}

// Sections provides a map sections with the unqualified section name as the index.
type Sections struct {
	Map map[string]Section
}

// Add sets map[key] = value.  key = an unqualified section name.
func (s *Sections) Add(key string, value Section) {
	s.Map[key] = value
}

// ProcessTemplate parses the content of the specified template file for the purpose of extracting section information.
func (s *Sections) ProcessTemplate(filename string) {
	currentState = LineStart
	var tpath TPath
	var scnr scanner.Scanner
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
	}
	scnr.Init(f)
	scnr.Filename = f.Name()
	scnr.Whitespace = 1<<':' | 1<<'-' | 1<<'\t' | 1<<'\n' | 1<<' ' // treat these as white space

	for tok := scnr.Scan(); tok != scanner.EOF; tok = scnr.Scan() {
		//		tokenError := TokenError{}

		token := scnr.TokenText()

		switch token {
		case "Insert_section":
			currentState = GotInsertSection
		case "End-Insert":
			currentState = GotEndSection
		case "Section":
			currentState = GotSection
		case "End-Section":
			currentState = GotEndSection
		case "Template":
			currentState = GotTemplateName
		default:
			switch currentState {
			case GotInsertSection:
				// Split on . to see if is fully qualified
				// Insert_section Selector_FestalApolytikion.af End-Insert
				// ?? left of first . = template name, subsequent = section.subsection
				//Sections with names ending in __ require a selection
				//Sections with names ending in _ have optional selections
				//Sections with names not ending in _ have no options
				// section path = template name.parentSection{0..n}.sectionName, e.g.
				// Insert_section MA._12_Exaposteilarion__.weekday__.exap1 End-Insert
			case GotEndInsert:
			case GotSection:
				section := s.Map[token]
				if len(section.Name) == 0 {
					section.Name = token
				}
				section.AddWhereDeclared(tpath)
				currentState = Neutral
			case GotEndSection:
			case GotTemplateName:
				tpath.Path = filename
				tpath.TemplateName = token
				currentState = Neutral
			}
		}
	}
}
