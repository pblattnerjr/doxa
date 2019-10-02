// Package css provides formatting for liturgical texts
package css

import (
	"github.com/liturgiko/doxa/pkg/utils/ltfile"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type SpanCss struct {
	Name   string // Text
	Color  string // font-Color: black
	Style  string // font-Style: italic, normal
	Weight string // font-Weight: normal, bold
}
func (s *SpanCss) Set(color Color, style Style, weight Weight) {
	s.Color = ColorMap[color].Name
	s.Style = StyleMap[style].Name
	s.Weight = WeightMap[weight].Name
	s.Name = CssClassName(color, style, weight)
}
// Returns a CssSpanArray of the specified size and sets color to BLACK, style to NORMALStyle, and weight to NORMALWeight
func NewCssSpanArray(size int) []SpanCss {
	var array []SpanCss
	for i := 0; i < size; i++ {
		array = append(array, NewSpanCss(CssClassName(BLACK, NORMALStyle, NORMALWeight)))
	}
	return array
}
// Returns a CssSpanArray of the indicated size and sets each span to the specified style
// The color will be BLACK and weight will be NORMALWeight
func NewCssSpanArrayStyle(size int, style Style) []SpanCss {
	var array []SpanCss
	for i := 0; i < size; i++ {
		array = append(array, NewSpanCss(CssClassName(BLACK, style, NORMALWeight)))
	}
	return array
}
// Returns a CssSpanArray of the indicated size and sets each span to the specified color
// The style will be NORMALStyle and the weight will be NORMALWeight
func NewCssSpanArrayColor(size int, color Color) []SpanCss {
	var array []SpanCss
	for i := 0; i < size; i++ {
		array = append(array, NewSpanCss(CssClassName(color, NORMALStyle, NORMALWeight)))
	}
	return array
}
// Returns a CssSpanArray of the indicated size and sets each span to the specified weight.
// The color will be BLACK and style will be NORMALStyle
func NewCssSpanArrayWeight(size int, weight Weight) []SpanCss {
	var array []SpanCss
	for i := 0; i < size; i++ {
		array = append(array, NewSpanCss(CssClassName(BLACK, NORMALStyle, weight)))
	}
	return array
}
// Returns a CssSpanArray of the specified size and sets color, style, and weight the same
func NewCssSpanArrayCSW(size int, c Color, s Style, w Weight) []SpanCss {
	var array []SpanCss
	for i := 0; i < size; i++ {
		array = append(array, NewSpanCss(CssClassName(c,s,w)))
	}
	return array
}
func NewSpanCss(name string) SpanCss {
	return SpanCss{name,"black", "normal", "normal"}
}
type Color int
type Style int
type Weight int
const (
	BLACK Color = iota
	RED
)
const (
	ITALIC Style = iota
	NORMALStyle
)
const (
	BOLD Weight = iota
	NORMALWeight
)
// maps color, style, and weight to a canonical css class name
func CssClassName(c Color, s Style, w Weight) string {
	sb := strings.Builder{}
	switch c {
	case RED: {sb.WriteString(ColorMap[RED].Abr)}
	default: {sb.WriteString(ColorMap[BLACK].Abr)}
	}
	switch s {
	case ITALIC: {sb.WriteString(StyleMap[ITALIC].Abr)}
	default: {sb.WriteString(StyleMap[NORMALStyle].Abr)}
	}
	switch w {
	case BOLD: {sb.WriteString(WeightMap[BOLD].Abr)}
	default: {sb.WriteString(WeightMap[NORMALWeight].Abr)}
	}
	return sb.String()
}
var Colors = []Color{BLACK, RED}
var Styles = []Style{ITALIC, NORMALStyle}
var Weights = []Weight{BOLD,NORMALWeight}

type elementData struct {
	Name string
	Abr string
}
var ColorMap = map[Color]elementData{
	BLACK: {"black", "Bk"},
	RED: {"red", "Rd"},
}
var StyleMap = map[Style]elementData{
	ITALIC: {"italic", "It"},
	NORMALStyle: {"normal", ""},
}
var WeightMap = map[Weight]elementData{
	BOLD: {"bold", "Bd"},
	NORMALWeight: {"normal",""},
}

const cssSpanTmpl = `
@font-face {
	font-family: "Arimo";
	src:
		url('/static/fonts/Arimo-Bold.ttf') format('truetype'),
		url('/static/fonts/Arimo-BoldItalic.ttf') format('truetype'),
		url('/static/fonts/Arimo-Italic.ttf') format('truetype'),
		url('/static/fonts/Arimo-Regular.ttf') format('truetype');
}

@font-face {
  font-family: 'ED Psaltica';
  src: url('/static/fonts/Psaltica.ttf') format('truetype');
  font-weight: normal;
  font-style: normal;
}
.container {
    position: absolute;
    left: -10px;
    height: 100%;
}
div.content {
	position: absolute;
	top: 20px;
	left: 5px;
}
html, body {
	font-family: "Arimo", sans-serif;
	font-size: 12pt;
    background-color: #FBF0D9;
    color: black; /* do not remove. Used in Javascript */
    height: 100%;
    width: 100%;
}
p {
	color: black;   /* night mode #FBF0D9 */
    text-align: left;
    font-family: "Arimo", sans-serif;
    padding-left: -50px;
    margin-bottom: 0em; /* to control spacing after para */
}
p.Dialog, p.DialogP {
	text-indent: 20px;
}
.ages-table-element {
    word-break: inherit;
    white-space: normal;
}
table {
  margin: 0 3pt;
  padding: 0;
  border-collapse: separate;
  border-spacing:5px;
  display: inline;
}
tr {
}
td {
  vertical-align: top;
  width: 50%;
  padding-left: 8px;
  padding-right: 8px;
}
td.leftCell {
  border-right: thin solid silver;
}
td.cellOneOfThree {
    vertical-align: top;
    width: 30vw;
	max-width: 30vw !important;
    padding-left: 8px;
    padding-right: 8px;
    border-right: thin solid silver;
}
td.cellTwoOfThree {
    vertical-align: top;
    width: 30vw;
    padding-left: 8px;
    padding-right: 8px;
    border-right: thin solid silver;
}
td.cellThreeOfThree {
    vertical-align: top;
    width: 30vw;
    padding-left: 8px;
    padding-right: 8px;
}
span.Error {
	color: red;
    font-weight: bold;
	border: 1px solid red;
}
{{range .}}
span.{{.Name}}{
  color: {{.Color}};
  font-style: {{.Style}};
  font-weight: {{.Weight}};
}{{end}}`

func WriteCss(path, filename string) error {
	var spans []SpanCss
	for _, color := range Colors {
		for _, style := range Styles {
			for _, weight := range Weights {
				var span SpanCss
				span.Name = CssClassName(color, style, weight)
				span.Color = ColorMap[color].Name
				span.Style = StyleMap[style].Name
				span.Weight = WeightMap[weight].Name
				spans = append(spans,span)
			}
		}
	}
	t := template.Must(template.New("spans").Parse(cssSpanTmpl))
	ltfile.CreateDirs(path)
	f, err := os.Create(filepath.Join(path, filename))
	if err != nil {
		log.Println("create file: ", err)
		return err
	}
	if err := t.Execute(f,spans); err != nil {panic(err)}
	return err
}
// ColorStyleWeigthSize
// {b|r}{i|n}{b|n}{one of the below}
//rnb = TextRedBold
//rnb25 = TextRedBoldEm25
func sizeToCss(size string) string {
	switch size {
	case "EM25": return "2.5em"
	case "EM20": return "2.0em"
	case "EM17": return "1.7em"
	case "EM14": return "1.4em"
	case "EM12": return "1.2em"
	case "EM10": return "1.0em"
	case "EM09": return "0.9em"
	case "EM08": return "0.8em"
	case "EM07": return "0.7em"
	case "EM05": return "0.5em"
	default: return "1.0em"
	}
}

func sizeCss() string {
	var sb strings.Builder
	sb.WriteString("span.EM25 {\n  font-size: 2.5em;\n}\n")
	sb.WriteString("span.EM20 {\n  font-size: 2.0em;\n}\n")
	sb.WriteString("span.EM17 {\n  font-size: 1.7em;\n}\n")
	sb.WriteString("span.EM14 {\n  font-size: 1.4em;\n}\n")
	sb.WriteString("span.EM12 {\n  font-size: 1.2em;\n}\n")
	sb.WriteString("span.EM10 {\n  font-size: 1.0em;\n}\n")
	sb.WriteString("span.EM09 {\n  font-size: 0.9em;\n}\n")
	sb.WriteString("span.EM08 {\n  font-size: 0.8em;\n}\n")
	sb.WriteString("span.EM07 {\n  font-size: 0.7em;\n}\n")
	sb.WriteString("span.EM05 {\n  font-size: 0.5em;\n}\n")
	return sb.String()
}

