package tags

type tag struct {
	Map map[string]string
}
func (b *tag) add(t string) {
	b.Map[t] = "." + t
}
func (b *tag) load(filename, extension string) {
}