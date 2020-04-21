package nlp 
import (
	"fmt"
	"strings"
	"text/scanner"
)
type Token struct {
	Position string
	Text string
}
const PUNCTUATION = `.,;?!·:`

type Frequency struct {
	Map map[string]int
	Total int
}

func (f *Frequency) add(s string) {
	f.Map[s] += 1
}
// Tokenizes the specified text and increments the count for each token
func (f *Frequency) Count(source, text string, nopunct bool, lang string) {
	for _, token := range Tokens(source, text, nopunct, lang) {
		f.add(token.Text)
	}
}
// Tokenizes the given text.
func Tokens(source, text string, nopunct bool, lang string) []Token {
	var s scanner.Scanner
	s.Init(strings.NewReader(text))
	s.Filename = source
	var tokens []Token
	var apostrophe = rune('\'')
	s.Whitespace = 1<<'-' | 1<<'\t' | 1<<'\n' | 1<<' ' | 1<<'\''// treat these as white space
	addApostrophe := false

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		if nopunct {
			i := strings.Index(PUNCTUATION,s.TokenText())
			if i > -1 {
				continue
			}
		}
		var t Token
		t.Position = fmt.Sprintf("%s",s.Position)
		t.Text = s.TokenText()
		if addApostrophe {
			t.Text = "'" + t.Text
			addApostrophe = false
		}
		switch lang {
		case "en":
			if s.Peek() == apostrophe {
				addApostrophe = true
			}
		case "gr":
			if s.Peek() == apostrophe {
				t.Text = t.Text + "'"
			}
		}
		tokens = append(tokens,t)
	}
	return tokens
}
