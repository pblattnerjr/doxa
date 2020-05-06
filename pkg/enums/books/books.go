// Package books provides an enum of Liturgical Books
package books

// Book is used to enumerate the Liturgical Books
type Book int
const (
	Eothina Book = iota
	Euchologion
	Heirmologion
	Horologion
	Katavasias
	Lectionary
	Menaion
	Octochechos
	Other
	Pentecostarion
	Psalter
	Triodion
)
func (b Book) Name() string {
	switch b {
	case Eothina: return "Eothina"
	case Euchologion: return "Euchologion"
	case Heirmologion: return "Heirmologion"
	case Horologion: return "Horologion"
	case Katavasias: return "Katavasias"
	case Lectionary: return "Lectionary"
	case Octochechos: return "Octochechos"
	case Other: return "Other"
	case Pentecostarion: return "Pentecostarion"
	case Psalter: return "Psalter"
	case Triodion: return "Triodion"
	default: return "unknown"
	}
}


