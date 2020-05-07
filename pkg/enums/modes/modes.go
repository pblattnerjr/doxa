// Package modes provides an enum of Liturgical Modes (aka Tones)
package modes

type Mode int
const (
	M1 Mode = iota + 1
	M2
	M3
	M4
	M5
	M6
	M7
	M8
)
func (m Mode) Name() string {
	switch m {
	case M1: return "Mode 1"
	case M2: return "Mode 2"
	case M3: return "Mode 3"
	case M4: return "Mode 4"
	case M5: return "Mode 5"
	case M6: return "Mode 6"
	case M7: return "Mode 7"
	case M8: return "Mode 8"
	default: return "unknown"
	}
}


