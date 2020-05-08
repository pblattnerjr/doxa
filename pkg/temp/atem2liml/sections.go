package atem2lml

import (
	"fmt"
	"github.com/liturgiko/doxa/pkg/utils/ltfile"
	"github.com/liturgiko/doxa/pkg/utils/stack"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/scanner"
)

type Section struct {
	Paths []string
}

func (s *Section) addPath(path string) {
	s.Paths = append(s.Paths, path)
}
func Index(dirIn, library string) (map[string]Section, error) {
	paths := stack.New()
	internalPaths := stack.New()

	var Map = map[string]Section{}

	files, err := ltfile.FileMatcher(dirIn, "atem", nil)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		parts := strings.Split(f, "a-templates/")
		if len(parts) == 2 {
			filename := filepath.Base(f)
			filename = filename[0:len(filename)-5]
			currentAGESPath := filename
			currentDoxaPath := library + filepath.Dir(parts[1]) + pathSeparator + filename
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
					currentDoxaPath = paths.Pop().(string)
					currentAGESPath = internalPaths.Pop().(string)
					currentState = Neutral
				default:
					switch currentState {
					case GotSection:
						var section Section
						ok := false
						if section, ok = Map[token]; ! ok {
							section = Section{}
						}
						section.Paths = append(section.Paths, currentDoxaPath)
						internalPaths.Push(currentAGESPath)
						currentAGESPath = currentAGESPath + "." + token
						Map[currentAGESPath] = section
						currentState = Neutral
						paths.Push(currentDoxaPath)
						currentDoxaPath = currentDoxaPath + pathSeparator + token
					}
				}
			}
		} else {
			fmt.Errorf("Should be using a-templates in path %s\n", f)
		}
	}
	return Map, nil
}
