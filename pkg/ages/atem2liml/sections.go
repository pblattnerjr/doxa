package atem2lml

import (
	"fmt"
	"github.com/emirpasic/gods/stacks/arraystack"
	"github.com/emirpasic/gods/trees/btree"
	"github.com/liturgiko/doxa/pkg/utils/ltfile"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/scanner"
)
// TemplateNode is used as a value in a tree, with the key being the name of the template.
// ID is the one that will be used by Doxa templates that use the liturgical markup language (lml).
// Path gives the path to the atem file.
// Sections is a tree index of the sections found in the template.
type TemplateNode struct {
	ID string
	Path string
	Sections *btree.Tree
}
// SectionNode holds the name of the template it occurs in and the path to the section.
// The TemplateName can be used to find the template in the Templates btree.
// The Path can be used as the Doxa path in a Doxa insert_section.
type SectionNode struct {
	TemplateName string
	Path string
}
// Index creates an index of all *.atem (AGES Template) files occurring recursively within the dirIn.
// It also creates an index of all sections in each template.
// The 1st btree returned is the templates index.  Its key is the template name and value is TemplateNode.
// The 2nd btree returned is the sections index.  Its key is the fully qualified path to the section.  Its value is the template name.
// So, if a key exists in the sections btree, the value can be used to retrieve the corresponding value
// in the templates btree.
// If dirIn does not exist, an error is returned and the tree will be empty.
func Index(dirIn, library string) (*btree.Tree, *btree.Tree, error) {
	paths := arraystack.New()
	internalPaths := arraystack.New()

	// tTree key = template name, value = TemplateNode
	tTree := btree.NewWithStringComparator(3) // order indicates max number of children per node
	// sTree key = template + "." + internal path to section, value = Template Name
	sTree := btree.NewWithStringComparator(3) // order indicates max number of children per node

	files, err := ltfile.FileMatcher(dirIn, "atem", nil)
	if err != nil {
		return tTree, sTree, err
	}
	for _, f := range files {
		var tNode TemplateNode
		tNode.Path = f
		tNode.Sections = btree.NewWithStringComparator(3)
		parts := strings.Split(f, "a-templates/")
		if len(parts) == 2 {
			filename := filepath.Base(f)
			filename = filename[0:len(filename)-5]
			currentAGESPath := filename
			currentDoxaPath := library + filepath.Dir(parts[1]) + pathSeparator + filename
			tNode.ID = currentDoxaPath
			currentState = LineStart
			// read file in as a string
			content, err := ioutil.ReadFile(f)
			if err != nil {
				fmt.Println(err.Error())
			}
			text := string(content)
			// Problem: text/scanner treats hyphen as a word separator.  Change to underscore to get around this.
			text = strings.ReplaceAll(text, "-", "_")

			var s scanner.Scanner
			s.Init(strings.NewReader(text))
			s.Whitespace = 1<<':' | 1<<'-' | 1<<'\t' | 1<<'\n' | 1<<' ' // treat these as white space

			for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
				token := s.TokenText()
				switch token {
				case "\r":
					continue
				case "Section":
					currentState = GotSection
				case "End_Section":
					if v, ok := paths.Pop(); ok {
						currentDoxaPath = v.(string)
					} else {
						fmt.Println("expected item from stack but stack has not values")
					}
					if v, ok := internalPaths.Pop(); ok {
						currentAGESPath = v.(string)
					} else {
						fmt.Println("expected item from stack but stack has not values")
					}
					currentState = Neutral
				default:
					switch currentState {
					case GotSection:
						internalPaths.Push(currentAGESPath)
						currentAGESPath = currentAGESPath + "." + token
						paths.Push(currentDoxaPath)
						currentDoxaPath = currentDoxaPath + pathSeparator + token
						var sNode SectionNode
						sNode.TemplateName = filename
						sNode.Path = currentDoxaPath
						sTree.Put(currentAGESPath, sNode)
						tNode.Sections.Put(currentAGESPath, currentDoxaPath)
						currentState = Neutral
					}
				}
				tTree.Put(filename, tNode)
			}
		} else {
			fmt.Errorf("Should be using a-templates in path %s\n", f)
		}
	}
	return tTree, sTree, nil
}
