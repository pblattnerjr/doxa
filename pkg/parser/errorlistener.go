package parser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type ParseError struct {
	TemplateID string
	Line int
	Column int
	Message string
}
// String returns an error formatted as line:column:message
func (e *ParseError) String() string {
	return fmt.Sprintf("%d:%d %s",e.Line, e.Column, e.Message)
}
// StringVerbose returns an error formatted as Template templateID Line line number Column column number: message
// e.g. Template a/b Line 1 Column 3: mismatched input '==' expecting '='
func (e *ParseError) StringVerbose() string {
	return fmt.Sprintf("Template %s Line %d Column %d: %s",e.TemplateID,e.Line, e.Column, e.Message)
}
type LMLErrorListener struct {
	antlr.ErrorListener
	TemplateID string
	Errors []ParseError
}
func NewLMLErrorListener(templateID string) *LMLErrorListener {
	l := new(LMLErrorListener)
	l.TemplateID = templateID
	return l
}
func (l *LMLErrorListener) AddError(templateID string, line, column int, msg string) {
	theError := ParseError{l.TemplateID, line,column, msg}
	l.Errors = append(l.Errors, theError)
}
func (l *LMLErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.AddError(l.TemplateID, line,column, msg)
}
func (l *LMLErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet){
//	fmt.Println("Unhandled ReportAmbiguity exception")
}
func (l *LMLErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet){
//	fmt.Println("Unhandled ReportAttemptingFullContext exception")
}
func (l *LMLErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet){
//	fmt.Println("Unhandled ReportContextSensitivity exception")
}

