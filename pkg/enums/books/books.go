//go:generate enumer -type=Book -json -text -yaml -sql
// 1) go get github.com/alvaroloes/enumer
// 2) in the enum subfolder for this enum: go generate

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


