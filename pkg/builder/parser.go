package builder

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"text/scanner"
)
type LookupMethod int
const (
	NID LookupMethod = iota // No ID
	RID // Relative ID
	SID // Specific ID
)
type Status int
const (
	NA Status = iota // No Status
	DRAFT
	FINAL
)
type Type int
const (
	BOOK Type = iota
	SERVICE
)

type TokenError struct {
	Token string
	Filename string
	Line   int
	Column int
	Error  error
}
func (t *TokenError) Set(s *scanner.Scanner, msg string) {
	currentState = ERROR
	t.Token = s.TokenText()
	t.Filename = s.Filename
	t.Line = s.Line
	t.Column = s.Column
	t.Error = errors.New(msg)

}
func (t *TokenError) String() string {
	return fmt.Sprintf("%s:%d:%d - %s - error: %s\n", t.Filename, t.Line, t.Column, t.Token, t.Error)
}
type Directive struct {
	Name LookupMethod
	Arg  string
}

type Span struct {
	Class string
	Id    string
	Value string
	Directives []Directive
	Error TokenError
}

func (s *Span) addDirective(d Directive) {
	s.Directives = append(s.Directives, d)
}

func newSpan() Span {
	var span Span
	span.Class = "Bk"
	return span
}
type Row struct {
	Cells []Cell
}
func (r *Row) addCell(c Cell) {
	r.Cells = append(r.Cells, c)
}

type Cell struct {
	Class       string
	Col         string
	Spans       []Span
}
func (c *Cell) addSpans(s Span, libraries []string) {
	var sb strings.Builder
	for _, library := range libraries {
		var span Span
		span.Class = s.Class

		sb.WriteString(library)
		sb.WriteString("~")
		sb.WriteString(s.Id)
		span.Id = sb.String()
		sb.Reset()

		span.Value = getValue(span.Id)
		span.Error = s.Error
		span.Directives = s.Directives
		c.Spans = append(c.Spans, span)
	}
}

// After parsing an lml file, the data will all be held here
type MetaTemplate struct {
	ID string
	Type Type
	Status Status
	Title string
	Day int
	Month int
	Year int
	Rows  []Row
}
func (m *MetaTemplate) addRow(r Row) {
	m.Rows = append(m.Rows, r)
}
var mt MetaTemplate

// ParseFiles parses all lml files matching the supplied regular expression
func ParseFiles(matching string) {

}
// ParseFile parses the specified lml file
func ParseFile(name string) {

}
// TODO: refactor so works just on content
func ParseTemplate(name, contents string) {
	libraries := []string{"gr_gr_cog","en_us_goa"}
	currentState = LineStart
	currentSpanClass := "Bk"
	lookupType := "SID"

	var row Row
	var cell Cell
	span := newSpan()

	var s scanner.Scanner
	f, err := os.Open(name)
	if err != nil {
		fmt.Println(err.Error())
	}
	s.Init(f)
	s.Filename = f.Name()
	s.Whitespace = 1<<':' | 1<<'-' | 1<<'\t' | 1<<'\n' | 1<<' '// treat these as white space
	lineNbr := 0

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		tokenError := TokenError{}

		if s.Position.Line > lineNbr {
			if lineNbr > 0 && len(cell.Class) > 0 {
				row.addCell(cell) // later this should be changed so we add the libraries and have a cell for each library
				mt.addRow(row)
				row = Row{}
				cell = Cell{}
				span = newSpan()
			}
			currentState = LineStart
			currentSpanClass = "Bk"
			lookupType = "SID"
			lineNbr = s.Position.Line
		}
		token := s.TokenText()

		switch token {
		case ".":
			switch currentState {
			case GotDotP, GotDotSpan:
				tokenError.Set(&s, "expected css class name after a dot")
			case GotRid, GotSid, GotNameTopicKey:
				currentState = GotDotSpan
			default:
				currentState = GotDotP
			}
		case "id":
			currentState = GotId
		case "nid":
			switch currentState {
			case GotDotP, GotDotSpan, GotNid, GotRid, GotSid:
				tokenError.Set(&s, "expected css class name or quoted string")
			default:
				currentState = GotNid
				lookupType = "NID"
				span.addDirective(Directive{Name: NID})
			}
		case "rid":
			switch currentState {
			case GotDotP, GotDotSpan, GotNid, GotRid, GotSid:
				tokenError.Set(&s, "expected quoted topic~key")
			default:
				currentState = GotRid
				lookupType = "RID"
				span.addDirective(Directive{Name: RID})
			}
		case "sid":
			switch currentState {
			case GotDotP, GotDotSpan, GotNid, GotRid, GotSid:
				tokenError.Set(&s, "expected quoted topic~key")
			default:
				currentState = GotSid
				lookupType = "SID"
				span.addDirective(Directive{Name: SID})
			}
		case "title":
			switch currentState {
			case LineStart:
				currentState = GotTitle
			default:
				tokenError.Set(&s, "keyword 'title' must start line")
			}
		case "type":
			switch currentState {
			case LineStart:
				currentState = GotType
			default:
				tokenError.Set(&s, "keyword 'type' must start line")
			}
		case "{", "}": // ignore
		default:
			quoted := strings.HasPrefix(s.TokenText(), "\"")
			switch currentState {
			case GotId:
				if quoted {
					mt.ID = trimQuotes(s.TokenText())
				} else {
					tokenError.Set(&s, "id must be quoted")
				}
			case GotDotP:
				if quoted { // css class name for p must not be quoted
					tokenError.Set(&s, "css class name must not be quoted")
				} else {
					currentState = GotNameP
					cell.Class = s.TokenText()
				}
			case GotDotSpan:
				if quoted { // css class name for span must not be quoted
					tokenError.Set(&s, "css class name must not be quoted")
				} else {
					currentState = GotNameSpan
					currentSpanClass = s.TokenText()
					span.Class = s.TokenText()
				}
			case GotNid, GotSid, GotRid:
				switch currentState {
				case GotNid:
					currentState = GotNameTopicKey
					span.Value = trimQuotes(s.TokenText())
					cell.addSpans(span, libraries)
					span = newSpan()
				default:
					if strings.Contains(s.TokenText(), "~") {
						currentState = GotNameTopicKey
						span.Id = trimQuotes(s.TokenText())
						cell.addSpans(span, libraries)
						span = newSpan()
					} else { // argument for sid or rid must have topic and key separated by tilde
						tokenError.Set(&s, "argument after rid or sid must be quoted topic~key")
					}
				}
			case GotTitle:
				mt.Title = strings.ToLower(s.TokenText())
			case GotType:
				theType := strings.ToLower(s.TokenText())
				switch theType {
				case "book": mt.Type = BOOK
				case "service": mt.Type = SERVICE
				default: tokenError.Set(&s, "argument after 'type' must be 'book' or 'service'")
				}
			default:
				if quoted {
					tokenError.Set(&s, "missing nid, sid, or rid before this")
				} else if token == "~" {
					tokenError.Set(&s, "missing quotes around topic~key")
				} else {
					tokenError.Set(&s, "unknown directive")
				}
			}
		}
		switch currentState {
		case GotNameTopicKey:
			fmt.Printf("%s: %s: %s, S: %s, L: %s\n", s.Position, s.TokenText(), currentState.stateName(), currentSpanClass, lookupType)
		case ERROR:
			fmt.Printf(tokenError.String())
		default:
			fmt.Printf("%s: %s: %s\n", s.Position, s.TokenText(), currentState.stateName())
		}
	}
}

var currentState state

type state int

const (
	LineStart state = iota
	GotColon
	GotCss
	GotDay
	GotDotP
	GotDotSpan
	GotDraft
	GotId
	GotMonth
	GotNameDirective
	GotNameP
	GotNameSpan
	GotNameTopicKey
	GotNid
	GotService
	GotRid
	GotSid
	GotStatus
	GotTitle
	GotType
	ERROR
	END
)
func (s state) stateName() string {
	switch s {
	case LineStart: return "LINE_START"
	case GotId: return "GOT_ID"
	case GotDotSpan: return "GOT_DOT_SPAN"
	case GotNameDirective: return "GOT_NAME_DIRECTIVE"
	case GotNameP: return "GOT_NAME_P"
	case GotDotP: return "GOT_DOT_P"
	case GotNameSpan: return "GOT_NAME_SPAN"
	case GotNameTopicKey: return "GOT_NAME_TOPIC_KEY"
	case GotNid: return "GOT_NID"
	case GotRid: return "GOT_RID"
	case GotSid: return "GOT_SID"
	case GotStatus: return "GOT_STATUS"
	case GotTitle: return "GOT_TITLE"
	case GotType: return "GOT_TYPE"
	case ERROR: return "ERROR"
	case END: return "END"
	default: return "UNKNOWN"
	}
}
func trimQuotes(s string) string {
	if len(s) >= 2 {
		if s[0] == '"' && s[len(s)-1] == '"' {
			return s[1 : len(s)-1]
		}
	}
	return s
}

// simulate database lookup
func getValue(id string) string {
	switch id {
	case "gr_gr_cog~actors~Deacon": return "ΔΙΑΚΟΝΟΣ:"
	case "en_us_goa~actors~Deacon": return "Deacon:"
	case "gr_gr_cog~prayers~res08": return "Εὐλόγησον, Δέσποτα."
	case "en_us_goa~prayers~res08": return "Master, give the blessing."
	case "gr_gr_cog~actors~Priest": return "ΙΕΡΕΥΣ"
	case "en_us_goa~actors~Priest": return "PRIEST"
	case "gr_gr_cog~prayers~enarxis02": return "Εὐλογημένη ἡ Βασιλεία τοῦ Πατρὸς καὶ τοῦ Υἱοῦ καὶ τοῦ Ἁγίου Πνεύματος, νῦν καὶ ἀεί, καὶ εἰς τοὺς αἰῶνας τῶν αἰώνων."
	case "en_us_goa~prayers~enarxis02": return "Blessed is the Kingdom of the Father and of the Son and of the Holy Spirit, now and forever and to the ages of ages."
	default: return "unknown id " + id
	}
}
/**
My idea:
push tokens onto stack with position info.
if you are in a tag line, when you hit eol,
  pop the stack until you get the tag, saving the topic-keys
  use the tag to get the formatting info and use it to create
     the output.
Note the a ...string is actually a []string. So, when
  you call the generic function have the topic-keys in a string slice

GetErrors
- missing topic or key
- topic-key does not exist in gr_gr_cog
- unknown tag
- missing opening brace
- missing closing brace for opening brace
- tag must be followed by at least one topic-key
*/


