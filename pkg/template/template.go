package template

import (
	"github.com/liturgiko/doxa/pkg/enums/calendarTypes"
	"github.com/liturgiko/doxa/pkg/enums/directiveTypes"
	"github.com/liturgiko/doxa/pkg/enums/idTypes"
	"github.com/liturgiko/doxa/pkg/enums/statuses"
	"github.com/liturgiko/doxa/pkg/enums/templateTypes"
	"github.com/liturgiko/doxa/pkg/ldp"
	"time"
)

// Abstract Liturgical Template
type ALT struct {
	ID string
	Type templateTypes.TemplateType
	Status statuses.Status
	Calendar calendarTypes.CalendarType
	Month, Day, Year int
	HtmlCss string
	PDF *PDF
	LDP ldp.LDP
}
// Set Liturgical Day Properties for specified month, day, and year
func (a *ALT) SetLDPYMD(month, day, year int, calendarType calendarTypes.CalendarType) error {
	l, err := ldp.NewLDPYMD(year, month, day, calendarType)
	if err != nil {
		return err
	}
	a.LDP = l
	return nil
}
// SetLDP for current year, and supplied month and day
func (a *ALT) SetLDPMD(month, day int, calendarType calendarTypes.CalendarType) error {
	l, err := ldp.NewLDPMD(month, day, calendarType)
	if err != nil {
		return err
	}
	a.LDP = l
	return nil
}
// SetLDP for template year, month, and day
func (a *ALT) SetLDP() error {
	var l ldp.LDP
	var err error
	if a.Year == 0 {
		l, err = ldp.NewLDPMD(a.Month, a.Day, a.Calendar)
		if err != nil {
			return err
		}
	} else {
		l, err = ldp.NewLDPYMD(a.Year, a.Month, a.Day, a.Calendar)
		if err != nil {
			return err
		}
	}
	a.LDP = l
	return nil
}

type PDF struct {
	CSS string
	PageNbr int
	Title string
	Headers []Header
	Footers []Footer
}
// AddHeader appends a header to the PDF struct's slice of Headers
func (p *PDF) AddHeader(header Header) {
	p.Headers = append(p.Headers, header)
}
// AddFooter appends a footer to the PDF struct's slice of Footers
func (p *PDF) AddFooter(footer Footer) {
	p.Footers = append(p.Footers, footer)
}
// Parity indicates whether a Header/Footer is even or odd or both
type Parity int
const (
	Both Parity = iota
	Even
	Odd
)
// Position designates slots in a header/footer. The slots are left, center, or right.
type Position int
const (
	Left Position = iota
	Center
	Right
)
// Header indicates parity (whether it is for even or odd pages or both)
// and the content of each of three slots: left, center, right.
// A slot can be empty.  To determine whether a slot has content, call the functions
// HasLeftSlot, HasCenterSlot, and HasRightSlot.
type Header struct {
	Parity Parity
	Left Slot
	Center Slot
	Right Slot
}
func (h *Header) AddLeftDirective(directive PDFDirective) {
	h.Left.Directives = append(h.Left.Directives, directive)
}
func (h *Header) AddCenterDirective(directive PDFDirective) {
	h.Center.Directives = append(h.Center.Directives, directive)
}
func (h *Header) AddRightDirective(directive PDFDirective) {
	h.Right.Directives = append(h.Right.Directives, directive)
}
func (h *Header) HasLeftSlot() bool {
	return len(h.Left.Directives) > 0
}
func (h *Header) HasCenterSlot() bool {
	return len(h.Center.Directives) > 0
}
func (h *Header) HasRightSlot() bool {
	return len(h.Right.Directives) > 0
}
// NewFooter returns a header with parity set to Both.
func NewHeader() *Header {
	var header = new(Header)
	header.Parity = Both
	return header
}
// NewHeaderEven returns a header with parity set to Even.
// That is, it is for even numbered pages.
func NewHeaderEven() *Header {
	var header = new(Header)
	header.Parity = Even
	return header
}
// NewHeaderOdd returns a header with parity set to Odd.
// That is, it is for odd numbered pages.
func NewHeaderOdd() *Header {
	var header = new(Header)
	header.Parity = Odd
	return header
}
// Slot contains information about a left, center, or right slot of a header or footer
type Slot struct {
	Directives []PDFDirective
}
func (s *Slot) AddDirective(directive PDFDirective) {
	s.Directives = append(s.Directives, directive)
}
// Footer indicates parity (whether it is for even or odd pages or both)
// and the content of each of three slots: left, center, right.
// A slot can be empty.  To determine whether a slot has content, call the functions
// HasLeftSlot, HasCenterSlot, and HasRightSlot.
type Footer struct {
	Parity Parity
	Left Slot
	Center Slot
	Right Slot
}
// NewFooter returns a footer with parity set to Both.
func NewFooter() *Footer {
	var footer = new(Footer)
	footer.Parity = Both
	return footer
}
// NewFooterEven returns a footer with parity set to Even.
// That is, it is for even numbered pages.
func NewFooterEven() *Footer {
	var footer = new(Footer)
	footer.Parity = Even
	return footer
}
// NewFooterOdd returns a footer with parity set to Odd.
// That is, it is for odd numbered pages.
func NewFooterOdd() *Footer {
	var footer = new(Footer)
	footer.Parity = Odd
	return footer
}
func (f *Footer) AddLeftDirective(directive PDFDirective) {
	f.Left.Directives = append(f.Left.Directives, directive)
}
func (f *Footer) AddCenterDirective(directive PDFDirective) {
	f.Center.Directives = append(f.Center.Directives, directive)
}
func (f *Footer) AddRightDirective(directive PDFDirective) {
	f.Right.Directives = append(f.Right.Directives, directive)
}
func (f *Footer) HasLeftSlot() bool {
	return len(f.Left.Directives) > 0
}
func (f *Footer) HasCenterSlot() bool {
	return len(f.Center.Directives) > 0
}
func (f *Footer) HasRightSlot() bool {
	return len(f.Right.Directives) > 0
}
// PDFDirective indicates what is to be inserted into a header or footer.
// The Class string holds the span CSS class to be applied.
// If the type = InsertDate, then the Time property is used.
// If the type = InsertLookup, then the Lookup property is used.
// If the type = InsertLiteral, then the Literal property is used.
// TODO: figure out a way to ensure the proper properties are used based on the Type
type PDFDirective struct {
	Type    directiveTypes.DirectiveType
	Class   string
	Lookup  Lookup
	Literal string
	Time    time.Time
}
func NewDateDirective(class string, time time.Time) *PDFDirective {
	var dir = new(PDFDirective)
	dir.Type = directiveTypes.InsertDate
	dir.Class = class
	dir.Time = time
	return dir
}
func NewPageNbrDirective(class string) *PDFDirective {
	var dir = new(PDFDirective)
	dir.Type = directiveTypes.InsertPageNbr
	dir.Class = class
	return dir
}
func NewLiteralDirective(class, literal string) *PDFDirective {
	var dir = new(PDFDirective)
	dir.Type = directiveTypes.InsertLiteral
	dir.Class = class
	dir.Literal = literal
	return dir
}
func NewLookupDirective(library int) *PDFDirective {
	var dir = new(PDFDirective)
	dir.Type = directiveTypes.InsertLookup
	var lookup = new(Lookup)
	lookup.Library = library
	dir.Lookup = *lookup
	return dir
}
func (d *PDFDirective) AddLookupTK(idType idTypes.IDType, class, topicKey string) error {
	var err error
	var lookupID = new(LookupTopicKey)
	lookupID.Type = idType
	lookupID.Class = class
	lookupID.TopicKey = topicKey
	d.Lookup.TopicKeys = append(d.Lookup.TopicKeys,*lookupID)
	return err
}
// Lookup provides information to do a database lookup to insert the result in a header/footer.
type Lookup struct {
	TopicKeys []LookupTopicKey
	Library   int
}
// LookupTopicKey indicates the type of lookup (RID or SID) and the Topic-Key to use and the CSS style class to use.
type LookupTopicKey struct {
	Type     idTypes.IDType
	Class string
	TopicKey string
	OverrideDay int
	OverrideMode int
}
/*
TopicKeySpan contains information for the formatting of a span that will contain the value from a database read using an ID.
Class is the CSS classname
Type is the type of Topic-Key (NID, SID, RID).  Technically NID is a literal, with no database lookup, but included for symmetry.
DataKey initially holds just the topic/key, but during generation becomes library/topic/key
Value is set during generation using the value from a database read
OverrideDay is used for a RID. Zero = no override. Otherwise is a value from 1-7, representing the days of the week
OverrideMode is used for a RID. Zero = no override.  Otherwise is a value from 1-8, representing the liturgical modes
 */
type TopicKeySpan struct {
	Class string
	Type idTypes.IDType
	DataKey string
	Value string
	OverrideDay int
	OverrideMode int
}
// AddLibrary prefixes the topic/key with the library so it is complete for a database lookup
func (tks *TopicKeySpan) AddLibrary(library string) {
	tks.DataKey = library + tks.DataKey
}
type Span struct {
	Class string
	TopicKeys []TopicKeySpan
	ChildSpans []Span
}
// AddTopic will create a TopicKeySpan and set the class to "kvp", and the DataKey to the topic/key.
// At generation time, the DataKey will be prefixed with the library so that an inspection of the HTML (for example) shows the ID used to get the value.
// The value will also be added at generation time.
func (s *Span) AddTopicKey(idType idTypes.IDType, topicKey string) {
	tks := new(TopicKeySpan)
	tks.Type = idType
	tks.Class = "kvp"
	tks.DataKey = topicKey
	s.TopicKeys = append(s.TopicKeys,*tks)
}
func (s *Span) AddChildSpan(span Span) {
	s.ChildSpans = append(s.ChildSpans,span)
}
type Paragraph struct {
	Class string
	Spans[]Span
	Version Span
}
func (p *Paragraph) AddSpan(span Span) {
	p.Spans = append(p.Spans,span)
}
// AddVersion creates a TopicKeySpan whose class = "versiondesignation" and an inner sid TopicKeySpan whose topic/key is "properties/version.designation"
// The version will display as an acronym for the library used during generation.
// For example, en_us_dedes will show [SD] as the version.
func (p *Paragraph) AddVersion(span Span) {
	vs := new(Span)
	vs.Class = "versiondesignation"
	vs.AddTopicKey(idTypes.SID, "properties/version.designation")
	p.Version = *vs
}
type Cell struct {
	Paragraph Paragraph
}
type Row struct {
	Class string
	Cells []Cell
	Version string
}
func (r *Row) AddCell(cell Cell) {
	r.Cells = append(r.Cells,cell)
}
