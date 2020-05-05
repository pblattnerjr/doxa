package atem2lml

import (
	"fmt"
	"os"
	"text/scanner"
)
type TPath struct {
	Path string
	TemplateName string
}
// Section provides a slice for which templates insert it and which name it.
// Because a section name is prefixed by a template name, it is possible for the
// same name to be declared by multiple templates.
type Section struct {
	SectionName string
	WhereInserted []TPath
	WhereNamed    []TPath
}
// AddInsert adds a template path to the list of where this section is inserted.
func (s *Section) AddInsert(path TPath) {
	s.WhereInserted = append(s.WhereInserted, path)
}
// AddWhereNamed adds a template path to the list of where this section is named.
// Multiple templates can have sections of the same name.
func (s *Section) AddWhereNamed(path TPath) {
	s.WhereNamed = append(s.WhereNamed, path)
}
// Sections provides a map sections with the section name as the index.
type Sections struct {
	Map map[string]Section
}
// Add sets map[key] = value.
func (s *Sections) Add(key string, value Section) {
	s.Map[key] = value
}

func (s *Sections) ProcessTemplate(filename string) {
	currentState = LineStart
	var currentSpanClass, lookupType string
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
		tokenError := TokenError{}

		token := scnr.TokenText()

		switch token {
		case "Insert_section": currentState = GotInsert
		case "End-Insert": currentState = GotEndSection
		case "Section": currentState = GotSection
		case "End-Section": currentState = GotEndSection
		case "Template": currentState = GotTemplateName
		default:
			switch currentState {
			case GotInsert:
			case GotEndInsert:
			case GotSection:
			case GotEndSection:
			case GotTemplateName:
				tpath.Path = filename
				tpath.TemplateName = token
			}
		}
		switch currentState {
		case GotInsert:
			fmt.Printf("%s: %s: %s, S: %s, L: %s\n", scnr.Position, scnr.TokenText(), currentState.stateName(), currentSpanClass, lookupType)
		case GotEndInsert:
			fmt.Printf("%s: %s: %s, S: %s, L: %s\n", scnr.Position, scnr.TokenText(), currentState.stateName(), currentSpanClass, lookupType)
		case GotSection:
			fmt.Printf("%s: %s: %s, S: %s, L: %s\n", scnr.Position, scnr.TokenText(), currentState.stateName(), currentSpanClass, lookupType)
		case GotEndSection:
			fmt.Printf("%s: %s: %s, S: %s, L: %s\n", scnr.Position, scnr.TokenText(), currentState.stateName(), currentSpanClass, lookupType)
		case ERROR:
			fmt.Printf(tokenError.String())
		default:
			fmt.Printf("%s: %s: %s\n", scnr.Position, scnr.TokenText(), currentState.stateName())
		}
	}
}
