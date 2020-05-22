//go:generate stringer -type=Status
// Package modes provides an enum of statuses for liturgical artifacts, e.g. translation, template
package statuses

type Status int
const (
	NA Status = iota
	Draft
	Review
	Final
)


