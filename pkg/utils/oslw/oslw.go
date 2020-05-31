// Functions for the processing of resources and templates
// in the OSLW (OCMC ShareLatex Liturgical Workbench).
// OSLW uses LaTeX to generated PDF files for up to three languages.
package oslw

import (
	"bufio"
	"database/sql"
	"errors"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/liturgiko/doxa/pkg/db/ltx2sql"
	"github.com/liturgiko/doxa/pkg/models"
	"github.com/liturgiko/doxa/pkg/ages/ares"
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

// Reads all OSLW resource files in specified directory
// and writes them to a database opened for the specified
// dbName
func Res2Sql(dir string, dbName string, logger *log.Logger) error {
	var err error
	var update = true // used to avoid updating record for gr_gr_cog if already exists

	// prepare the database
	var db *sql.DB
	mapper := &ltx2sql.LtxMapper{}
	db, err = sql.Open("sqlite3", dbName)
	if err!= nil {
		return err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return err
	}
	_, err = db.Exec(ltx2sql.SQLCreateTable)
	if err!= nil {
		return err
	}
	mapper.DB = db

	// process the files
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
				var ltx models.Ltx
				ltx.ID = ltstring.ToId(
					lineParts.Language,
					lineParts.Country,
					lineParts.Realm,
					lineParts.Topic,
					lineParts.Key,
				)
				ltx.Topic = lineParts.Topic
				ltx.Key = lineParts.Key
				update = true
				if strings.HasPrefix(ltx.ID, "gr_gr_cog") {
					_, err = mapper.ReadById(ltx.ID)
					update = err != nil
				}
				if update {
					err = mapper.Merge(&ltx)
				}
				if err != nil {
					logger.Println(err.Error())
				}
			} else if strings.HasPrefix(line,"\\itId") {
				lineParts, err = ParseOslwResource(line, "")
				if err != nil {
					logger.Println(err.Error())
				}
			} else {
				lineParts.Value = line
			}
		}
		fileIn.Close()
	}
	return err
}