package template

import (
	"encoding/json"
	"fmt"
	"github.com/liturgiko/doxa/pkg/enums/statuses"
	"github.com/liturgiko/doxa/pkg/enums/templateTypes"
	"testing"
)

func TestALT(t *testing.T) {
	alt := new(ALT)
	alt.ID = "x/y"
	alt.Type = templateTypes.Service
	alt.Status = statuses.Draft
	alt.HtmlCss = "ages.html.css"
	pdf := new(PDF)
	pdf.Title = "I am the title"
	pdf.CSS = "ages.pdf.css"
	pdf.PageNbr = 1
	alt.PDF = pdf
	headerEven := NewHeaderEven()
	headerEven.AddCenterDirective(*NewDateDirective(".it"))
	pdf.AddHeader(*headerEven)
	headerOdd := NewHeaderOdd()
	headerOdd.AddCenterDirective(*NewPageNbrDirective(".it"))
	pdf.AddHeader(*headerOdd)
	footerEven := NewFooterEven()
	footerEven.AddCenterDirective(*NewDateDirective(".it"))
	pdf.AddFooter(*footerEven)
	footerOdd := NewFooterOdd()
	footerOdd.AddCenterDirective(*NewPageNbrDirective(".it"))
	pdf.AddFooter(*footerOdd)
	j, err := json.MarshalIndent(alt, "", "  m")
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(string(j))
}