//go:generate enumer -type=Position -json -text -yaml -sql
// 1) go get github.com/alvaroloes/enumer
// 2) in the enum subfolder for this enum: go generate

// Package modes provides an enum of header/footer slot positions (Left, Center, Right)
package positions

type Position int
const (
	Left Position = iota
	Center
	Right
)


