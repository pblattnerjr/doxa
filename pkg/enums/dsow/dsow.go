//go:generate enumer -type=DOW -json -text -yaml -sql
// 1) go get github.com/alvaroloes/enumer
// 2) in the enum subfolder for this enum: go generate

// Package dsows provides an enum of the days of the week
package dsow

type DOW int
const (
	D1 DOW = iota + 1
	D2
	D3
	D4
	D5
	D6
	D7
)
func (d DOW) Name() string {
	switch d {
	case D1: return "Sunday"
	case D2: return "Monday"
	case D3: return "Tuesday"
	case D4: return "Wednesday"
	case D5: return "Thursday"
	case D6: return "Friday"
	case D7: return "Saturday"
	default: return "unknown"
	}
}

