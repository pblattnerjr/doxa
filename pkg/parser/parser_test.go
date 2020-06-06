package parser

import (
	"encoding/json"
	"fmt"
	"github.com/liturgiko/doxa/pkg/enums/statuses"
	"github.com/liturgiko/doxa/pkg/enums/templateTypes"
	"os"
	"testing"
)

var dbPath = os.Getenv("dbPath")

func init() {
	// open the database
}
func TestLML_Errors(t *testing.T) {
	input := `ID == "xy"
Type = "services"
Status = "drafty"
Month = 66
Day = 33
if BadID == 1 { insert "actors/Priest"" }`
	lml, err := NewLMLParser("a/b",input, dbPath)
	if err != nil {
		t.Errorf("could not open database %s", dbPath)
	}
	for _, error := range lml.WalkTemplate() {
		fmt.Println(error.StringVerbose())
	}
}
func TestTemplate(t *testing.T) {
	input := `
ID = "ages/Dated-Services/m01/d06/se.m01.d06.li" // used as id in database
Type = "service" // types are block, book, or service. 
Status = "draft" // status values are na, draft, review, or final. 
Calendar = "Gregorian" // if omitted is set to Gregorian.  Can be set to Julian.
HtmlCss = "ages/css/html.css"
PdfCss = "ages/css/pdf.css"
Title = "Divine Liturgy" // optional. Used for searching db for template
Month = 1 // required if template type is a service
Day = 6  // required if template type is a service
Year = 2020 // optional for a service. If year is omitted or set to zero, the current year will be used.
PageHeader = center @Lookup sid "template.titles/ve.pdf.header" rid "da/daVE.OnTheEveningBefore" lang 1
PageFooter = left @PageNbr right @Date // You can use one or more of left, center, and right.
SetPageNumber = 1 
{
// 1a
//p.actor sid "actors/Priest"
// 1b
//p.actor sid "actors/Priest" sid "rubrical/InALowVoice"
// 2a
//p.dialog span.it sid "prayers/res04p"
// 2b
//p.dialog span.it sid "prayers/res04p" span.rubric sid "rubrical/Thrice"
// 3
//p.actor span.rditbd sid "actors/Deacon" (span.rubric sid "rubrical/InALowVoice") span.bk nid "But, loud enough to be heard."
// 4
//p.actor sid "actors/Deacon" (span.rubric sid "rubrical/InALowVoice") span.bk nid "But, loud enough to be heard."
// 5
//p.actor sid "actors/Deacon" (span.rubric sid "rubrical/InALowVoice") nid "But, loud enough to be heard."
// 6
//p.actor sid "actors/Deacon" (span.rubric sid "rubrical/InALowVoice" span.rubric sid "rubrical/Thrice") nid "But, loud enough to be heard."
// 7
//p.actor span.it sid "actors/Deacon" (span.rubric sid "rubrical/InALowVoice" span.rubric sid "rubrical/Thrice") nid "But, loud enough to be heard."
// 8
//p.actor span.it sid "actors/Deacon" (span.rubric sid "rubrical/InALowVoice"  sid "rubrical/Thrice") nid "But, loud enough to be heard."
// 9
p.actor span.it sid "actors/Deacon" (span.rubric sid "rubrical/InALowVoice" ( span.bl sid "rubrical/Thrice" ) ) nid "But, loud enough to be heard."


 // p.actor sid "actors/Priest"
  //p.dialog sid "eu.lichrysbasil/euLI.Key0109.text"
 // p.dialog span.it sid "prayers/res04p"
  //p.hymn rid "oc.*/ocVE.ApolTheotokionVM.text" @Mode 1 @Day 1 @Ver // @Mode and @Day override the mode and day for the service date.
  //p.actor span.rditbd sid "actors/Deacon" span.rubric sid "rubrical/InALowVoice" span.bk nid "But, loud enough to be heard."
  // p.actor span.rditbd sid "actors/Deacon" (span.rubric sid "rubrical/InALowVoice") span.bk nid "But, loud enough to be heard."
}`
	lml, err := NewLMLParser("a/b",input, dbPath)
	if err != nil {
		t.Errorf("could not open database %s", dbPath)
	}
	for _, error := range lml.WalkTemplate() {
		t.Error(error.StringVerbose())
	}
	if lml.Listener.ALT.ID != "ages/Dated-Services/m01/d06/se.m01.d06.li" {
		t.Errorf(fmt.Sprintf("got ID = %s, expected %s", lml.Listener.ALT.ID, "ages/Dated-Services/m01/d06/se.m01.d06.li"))
	}
	if lml.Listener.ALT.Type != templateTypes.Service {
		t.Errorf(fmt.Sprintf("got Type = %s, expected %s", lml.Listener.ALT.Type.String(), templateTypes.Service))
	}
	if lml.Listener.ALT.Status != statuses.Draft {
		t.Errorf(fmt.Sprintf("got Status = %s, expected %s", lml.Listener.ALT.Status.String(), statuses.Draft))
	}
	if lml.Listener.ALT.HtmlCss != "ages/css/html.css" {
		t.Errorf(fmt.Sprintf("got HtmlCss = %s, expected %s", lml.Listener.ALT.HtmlCss, "ages/css/html.css"))
	}
	if lml.Listener.ALT.PDF.CSS != "ages/css/pdf.css" {
		t.Errorf(fmt.Sprintf("got PdfCss = %s, expected %s", lml.Listener.ALT.PDF.CSS, "ages/css/pdf.css"))
	}
	if lml.Listener.ALT.PDF.Title != "Divine Liturgy" {
		t.Errorf(fmt.Sprintf("got Title = %s, expected %s", lml.Listener.ALT.PDF.Title, "Divine Liturgy"))
	}
	if lml.Listener.ALT.Month != 1 {
		t.Errorf(fmt.Sprintf("got Month = %d, expected %d", lml.Listener.ALT.Month, 1))
	}
	if lml.Listener.ALT.Day != 6 {
		t.Errorf(fmt.Sprintf("got Day = %d, expected %d", lml.Listener.ALT.Day, 6))
	}
	if lml.Listener.ALT.Year != 2020 {
		t.Errorf(fmt.Sprintf("got Year = %d, expected %d", lml.Listener.ALT.Year, 2020))
	}
	if lml.Listener.ALT.PDF.PageNbr != 1 {
		t.Errorf(fmt.Sprintf("got SetPageNumber = %d, expected %d", lml.Listener.ALT.PDF.PageNbr, 1))
	}
	if lml.Listener.ALT.LDP.TheDay.Year() != 2020 {
		t.Errorf(fmt.Sprintf("got LDP year = %d, expected %d", lml.Listener.ALT.LDP.TheDay.Year(), 2020))
	}
	if lml.Listener.ALT.LDP.TheDay.Month() != 1 {
		t.Errorf(fmt.Sprintf("got LDP month = %d, expected %d", lml.Listener.ALT.LDP.TheDay.Month(), 1))
	}
	if lml.Listener.ALT.LDP.TheDay.Day() != 6 {
		t.Errorf(fmt.Sprintf("got LDP day = %d, expected %d", lml.Listener.ALT.LDP.TheDay.Day(), 6))
	}
	json, err := json.MarshalIndent(lml.Listener.ALT, "", " ")
	fmt.Println(string(json))
}
func TestLML_Tokens(t *testing.T) {
	input := `ID = "x/y"`
	lml, err := NewLMLParser("x/y", input, dbPath)
	if err != nil  {
		t.Errorf("could not open database %s", dbPath)
	}
	for _, t := range lml.Tokens() {
		fmt.Printf("%d:%d %s (%q)\n", t.GetLine(), t.GetColumn(),
			lml.Lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
	for _, e := range lml.ErrorListener.Errors {
		fmt.Println(e)
	}
}
