//go:generate enumer -type=Status -json -text -yaml -sql
// 1) go get github.com/alvaroloes/enumer
// 2) in the enum subfolder for this enum: go generate

// Package modes provides an enum of statuses for liturgical artifacts, e.g. translation, template
package statuses

type Status int
const (
	NA Status = iota
	Draft
	Review
	Final
)


