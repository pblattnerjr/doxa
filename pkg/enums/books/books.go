//go:generate stringer -type=Book
//Package books provides an enum of Liturgical Books
package books

// Book is used to enumerate the Liturgical Books
type Book int
const (
	Eothina Book = iota
	Euchologion
	Heirmologion
	Horologion
	Katavasias
	Lectionaryd
	Menaion
	Octochechos
	Other
	Pentecostarion
	Psalter
	Triodion
)


