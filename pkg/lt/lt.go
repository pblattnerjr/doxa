package lt

/**
  TODO: figure out how to pass custom span formatting.  One approach is "rib~actors~Priest", where we split on ~ and if length > 2, we know the first part is the span format
  TODO: test via a doxago serve command
  TODO: create a LiturgicalDay (ldp) package
*/
import (
	"bytes"
	"errors"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jmoiron/sqlx"
	"github.com/liturgiko/doxa/pkg/css"
	"github.com/liturgiko/doxa/pkg/ldp"
	"github.com/liturgiko/doxa/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

var TemplateDir string = "/Users/mac002/doxago/templates"
var bootstrap *template.Template

type docProps struct {
	Ldp   ldp.LDP
	error string
}

var DocProps docProps

type Span struct {
	Class string
	Id    string
	Value string
	Colon bool
	Error bool
}
type Row struct {
	Cells []Cell
}
type Cell struct {
	Class       string
	Parentheses bool
	Col         string
	Spans       []Span
}

const (
	RIGHT    = "rightCell"
	CENTER   = "centerCell"
	LEFT     = "leftCell"
	CELL1Of3 = "cellOneOfThree"
	CELL2Of3 = "cellTwoOfThree"
	CELL3Of3 = "cellThreeOfThree"
)

type Data struct {
	Title string
	Rows  []Row
}

var Table Data

type Command string

var Domains []string

// The following commands are used in liturgical templates to set properties for the
// document to be generated.

// Sets the date for the generation of a document.
// If the values are all equal to zero, the date is set to today's date.
// If the year is 0 and month and day are > 0, the current year is used.
func (p Command) SetDate(year, month, day int) {
	var theLdp ldp.LDP
	var err error
	if year == 0 && month == 0 && day == 0 {
		theLdp, err = ldp.NewLDP()
	} else {
		if year == 0 && month > 0 && day > 0 {
			theLdp, err = ldp.NewLDPMD(month, day)
		} else if year > 1583 && month > 0 && day > 0 {
			theLdp, err = ldp.NewLDPYMD(year, month, day)
		} else {
			err = errors.New("year must be 0 or > 1583, month > 0, and day > 0")
		}
	}
	DocProps.Ldp = theLdp
	if err != nil {
		DocProps.error = err.Error()
	} else {
		DocProps.error = ""
	}
}

// This is the generic version of template commands used to format liturgical texts.
func (p Command) Generic(class string, a []string, f []css.SpanCss) string {
	var row Row
	for d, domain := range Domains {
		var sb strings.Builder
		var cell Cell
		cell.Class = class
		cell.Parentheses = strings.HasSuffix(class, "P")
		switch len(Domains) {
		case 1:
			cell.Col = LEFT
		case 2:
			switch d {
			case 0:
				cell.Col = LEFT
			case 1:
				cell.Col = RIGHT
			}
		case 3:
			switch d {
			case 0:
				cell.Col = CELL1Of3
			case 1:
				cell.Col = CELL2Of3
			case 2:
				cell.Col = CELL3Of3
			}
		default:
			cell.Col = RIGHT
		}
		for i, tk := range a {
			// if there i more than one span, we need a space between the values
			if i > 0 {
				sb.WriteString(" ")
			}
			id := domain + "~" + tk
			hasError := false
			ltext, err := GetRecord(id)
			if err != nil {
				hasError = true
				msg := err.Error()
				if strings.HasPrefix(msg, "sql: no rows") {
					ltext.Value = fmt.Sprintf("No record exists in database for ID %s.", id)
				} else {
					ltext.Value = fmt.Sprintf("Bad ID %s. %s", id, err.Error())
				}
			}
			sb.WriteString(ltext.Value)

			// if the cell class is an actor, we want a colon after the actor, e.g. PRIEST:
			if strings.HasPrefix(cell.Class, "Actor") && len(class) > 5 && i == 0 {
				sb.WriteString(":")
			}
			var span Span
			// if there was an error retrieving the record, we will use the span.Error case to
			// format it so the user can easily see it.
			if hasError {
				span.Class = "Error"
			} else {
				span.Class = f[i].Name
			}
			span.Id = id
			span.Value = sb.String()
			sb.Reset()
			cell.Spans = append(cell.Spans, span)
		}
		row.Cells = append(row.Cells, cell)
	}
	Table.Rows = append(Table.Rows, row)
	// we have to return a string in order to make this work
	// but an empty string will do the trick.
	return ""
}

// Liturgical template command
func (p Command) Actor(a ...string) string {
	// TODO: this should only accept a single a
	fmtArray := css.NewCssSpanArray(len(a))
	fmtArray[0].Set(css.RED, css.NORMALStyle, css.NORMALWeight)
	return p.Generic("Actor", a, fmtArray)
}

// Liturgical template command
func (p Command) Dialog(a ...string) string {
	fmtArray := css.NewCssSpanArray(len(a))
	return p.Generic("Dialog", a, fmtArray)
}

// Liturgical template command
func (p Command) DialogP(a ...string) string {
	fmtArray := css.NewCssSpanArrayStyle(len(a), css.ITALIC)
	return p.Generic("DialogP", a, fmtArray)
}

// Liturgical template command
func (p Command) ActorDialog(a ...string) string {
	fmtArray := css.NewCssSpanArray(len(a))
	fmtArray[0].Set(css.RED, css.NORMALStyle, css.NORMALWeight)
	return p.Generic("ActorDialog", a, fmtArray)
}

// Liturgical template command
func (p Command) Insert(a ...string) string {
	tmpl, err := template.ParseFiles(TemplateDir + "/" + a[0] + ".gohtml")
	if err != nil {
		log.Fatalf("Parse: %v", err)
	}
	var rows bytes.Buffer
	tmpl.Execute(&rows, Command("insert"))
	return ""
}

// Liturgical template command to temporarily override
// the movable cycle day.
func (p Command) SetMCDay(d int) {
	DocProps.Ldp.OverrideMovableCycleDay(d)
}

// stores the Ltext records that have already been retrieved from the database so we do not do another call to the db.
var Retrieved map[string]models.Ltext

// get the record with the specified id.  If it has been already retrieved from the database
// we will get it from the retrieved map.  Otherwise, we will read it from the database.
func GetRecord(id string) (models.Ltext, error) {
	if ltext, ok := Retrieved[id]; ok {
		return ltext, nil
	} else {
		ltext = models.Ltext{}
		err := Db.Get(&ltext, "SELECT * FROM ltext WHERE id=$1", id)
		return ltext, err
	}
}

var Db *sqlx.DB

func Serve(port, home string) {
	var err error
	// open the database
	Db, err = sqlx.Open("sqlite3", filepath.Join(home, "data", "sql", "liturgical.db"))
	if err != nil {
		panic(err)
	}
	// set up the domains we will process
	Domains = append(Domains, "gr_gr_cog")
	Domains = append(Domains, "en_us_dedes")
	//	Domains = append(Domains, "gr_gr_cog")

	bootstrap, err = template.ParseGlob(filepath.Join(home, "templates", "layout", "*.gohtml"))
	if err != nil {
		panic(err)
	}

	tmpl, err := template.ParseFiles(filepath.Join(home, "templates", "rows", "eu.li.chrysostom.gohtml"))
	if err != nil {
		log.Fatalf("Parse: %v", err)
	}
	// rows is a dummy variable.
	var rows bytes.Buffer
	tmpl.Execute(&rows, Command("rows"))
	Table.Title = "Divine Liturgy"

	fs := http.FileServer(http.Dir(filepath.Join(home, "http", "static")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handler)

	http.ListenAndServe(":"+port, nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	bootstrap.ExecuteTemplate(w, "bootstrap", Table)
}

