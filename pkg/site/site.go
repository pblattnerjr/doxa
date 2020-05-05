package site

import (
	"github.com/jmoiron/sqlx"
	"github.com/liturgiko/doxa/pkg/ldp"
	"github.com/liturgiko/doxa/pkg/utils/ltfile"
)

var Db *sqlx.DB
var TemplateDir string

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
var Domains []string

// For each domain, generate files of specified types whose names match one of the patterns
func Build(templatesDir string,
	dbPath string, // path to the sqlite database
	siteDir string, // path to the website directory
	patterns []string, // regular expressions to match template filenames to use for generation
	extension string, // template file extension to look for, without the period
	domains []string, // which domains to generate for
) error { // types of files to generate
	templates, err := ltfile.FileMatcher(templatesDir, extension, patterns)
	if err != nil {
		return err
	}
	TemplateDir = templatesDir
	for _, template := range templates {
		err = GenerateFromTemplate(templatesDir, dbPath, template, siteDir, domains)
		if err != nil {
			return err
		}
	}
	return err
}
func GenerateFromTemplate(templatesDir string,
	dbPath string,
	docTemplatePath string,
	outputPath string,
	domains []string) error {

	var err error
	// open the database
	Db, err = sqlx.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	// set up the domains we will process
	Domains = domains

	//doc, err = template.ParseGlob(filepath.Join(templatesDir, "layout", "*.gohtml"))
	//if err != nil {
	//	return err
	//}

	//tmpl, err := template.ParseFiles(docTemplatePath)
	//if err != nil {
	//	return err
	//}
	// rows is a dummy variable.
//	var rows bytes.Buffer
//	tmpl.Execute(&rows, Command("rows"))
	Table.Title = "Divine Liturgy"
// 	f, err := os.Create(filepath.Join(outputPath, "index.html"))
//	if err != nil {
//		return err
//	}
//	doc.ExecuteTemplate(f, "doc", Table)
	return err
}



