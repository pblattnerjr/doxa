//go:generate enumer -type=DirectiveType -json -text -yaml -sql
// 1) go get github.com/alvaroloes/enumer
// 2) in the enum subfolder for this enum: go generate

// Package modes provides an enum of directive types
package directiveTypes

type DirectiveType int
const (
	InsertDate DirectiveType = iota
	InsertLiteral
	InsertLookup
	InsertPageNbr
	InsertVersion
)


