//go:generate enumer -type=TemplateType -json -text -yaml -sql
// 1) go get github.com/alvaroloes/enumer
// 2) in the enum subfolder for this enum: go generate
package templateTypes

type TemplateType int
const (
	Block TemplateType = iota
	Book
	Service
)
