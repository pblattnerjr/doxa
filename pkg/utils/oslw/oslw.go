// Functions for the processing of resources and templates
// in the OSLW (OCMC ShareLatex Liturgical Workbench).
// OSLW uses LaTeX to generated PDF files for up to three languages.
package oslw

import (
	"bufio"
	"errors"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jmoiron/sqlx"
	"github.com/liturgiko/doxa/pkg/db/lsql"
	"github.com/liturgiko/doxa/pkg/models"
	"github.com/liturgiko/doxa/pkg/utils/ares"
	"github.com/liturgiko/doxa/pkg/utils/ltfile"
	"github.com/liturgiko/doxa/pkg/utils/ltstring"
	"log"
	"os"
	"strings"
)

/**
* resources
    en_uk_lash
      res.titles.tex
 */
/*
\itId{en}{uk}{lash}{actors}{Bishop}{
Bishop
}%
 */
func ParseOslwResource(id, value string) (ares.LineParts, error) {
	var result ares.LineParts
	var err error
	parts := strings.Split(id, "Id{")
	if len(parts) > 1 {
		parts = strings.Split(parts[1], "}{")
		if len(parts) == 6 {
			result.Language = parts[0]
			result.Country = parts[1]
			result.Realm = parts[2]
			result.Topic = parts[3]
			result.Key = parts[4]
			result.Value = value
		} else {
			err = errors.New("invalid OSLW ID format")
		}
	} else {
		err = errors.New("invalid OSLW ID format")
	}
	return result, err
}

// Reads all resource files in specified directory
// and writes them to a database opened for the specified
// dbName
func LoadOslwResources(dir string, dbName string) error {
	var err error

	// open the database
	var db *sqlx.DB
	db, err = sqlx.Open("sqlite3", dbName)
	if err != nil {
		return err
	}
	err = lsql.CreateSchema(models.LtextSchema, dbName)
	if err != nil {
		return err
	}

	ext := "tex"
	patterns := []string{
		"res.*",
	}
	filenames, err := ltfile.FileMatcher(dir, ext, patterns)
	if err != nil {
		return err
	}
	for _, filename := range filenames {
		fileIn, err := os.Open(filename)
		if err != nil {
			fileIn.Close()
			return err
		}
		scanner := bufio.NewScanner(fileIn)
		var lineParts ares.LineParts

		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line == "}%" {
				// write to db
				var ltext models.Ltext
				ltext.ID = ltstring.ToId(
					lineParts.Language,
					lineParts.Country,
					lineParts.Realm,
					lineParts.Topic,
					lineParts.Key,
				)
				ltext.Topic = lineParts.Topic
				ltext.Key = lineParts.Key
				err = ltext.CreateRecord(db)
				if err != nil {
					log.Println(err.Error())
				}
			} else if strings.HasPrefix(line,"\\itId") {
				lineParts, err = ParseOslwResource(line, "")
				if err != nil {
					log.Println(err.Error())
				}
			} else {
				lineParts.Value = line
			}
		}
		fileIn.Close()
	}
	return err
}