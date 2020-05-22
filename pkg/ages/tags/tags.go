package tags

import (
	"github.com/liturgiko/doxa/pkg/utils/ltfile"
	"strings"
	"text/scanner"
)

type Tag struct {
	Map map[string]string
}
func (b *Tag) add(key, value string) {
	b.Map[key] = "." + value
}
func (b *Tag) Load(dir, filename string) error {
	b.Map = make(map[string]string)
	patterns := []string{filename}
	files, err := ltfile.FileMatcher(dir, "ares", patterns)
	if err != nil {
		return err
	}
	for _, file := range files {
		lines, err :=  ltfile.GetFileLines(file)
		if err != nil {return err}
		for _, line := range lines {
			// skip over the topic name
			if strings.HasPrefix(line, "A_Resource_Whose_Name") {
				continue
			}
			var s scanner.Scanner
			s.Init(strings.NewReader(line))
			s.Whitespace = 1<<':' | 1<<'-' | 1<<'\t' | 1<<' ' // treat these as white space

			currentState := NOP
			var key string

			for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
				token := s.TokenText()

				switch token {
				case "=":
					{
					currentState = GOTEQUAL
					}
				default:
					switch currentState {
					case GOTEQUAL:
						{
							parts := strings.Split(token,",")
							if len(parts) == 2 {
								value := parts[1]
								b.add(key, value[:len(value)-1])
							} else {
								b.add(key, token)
							}
						}
					default:
						key = token
					}
				}
			}
		}
	}
	return nil
}
const (
	NOP = iota
	GOTEQUAL
)