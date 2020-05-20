//go:generate stringer -type=Status
// Package modes provides an enum of statuses for liturgical artifacts, e.g. translation, template
package status

type Status int
const (
	na Status = iota
	draft
	review
	final
)


