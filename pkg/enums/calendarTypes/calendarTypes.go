//go:generate enumer -type=CalendarType -json -text -yaml -sql
// 1) go get github.com/alvaroloes/enumer
// 2) in the enum subfolder for this enum: go generate

// Package modes provides an enum of calendar types for liturgical services
package calendarTypes

type CalendarType int
const (
	Gregorian CalendarType = iota
	Julian
)


