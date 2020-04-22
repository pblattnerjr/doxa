package lsql

// Package lsql (liturgical sqlite) provides models and functions for the loading of
// liturgical text into sqlite databases.

// TODO: 1) write sql more quickly.  The commented out code creates a sql file to load.  There are some errors when it loads.  Need to figure out why.

import (
	"errors"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jmoiron/sqlx"
	"github.com/liturgiko/doxa/pkg/mappers"
	"github.com/liturgiko/doxa/pkg/models"
	"github.com/liturgiko/doxa/pkg/utils/repos"
	"log"
	"os"
	//	"path/filepath"
	"strings"
	//	"text/template"
)


// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Merge a schema in a Sqlite3 database opened using a supplied path (dbname)
func CreateSchema(schema string, dbname string) error {
	db, err := sqlx.Connect("sqlite3", dbname)
	if err != nil {
		return err
	}
	_, err = db.Exec(schema)
	return err
}

// Using the repo urls loads the repos into memory from Github, then writes the lines from each file into the database.
func AresGithub2Sqlite(
	urls []string,
	dbname string,
	printProgress bool,
	logger *log.Logger,
) error {
	var err error
	// delete existing database
	if FileExists(dbname) {
		os.Remove(dbname)
	}
	// set up the schema
	err = CreateSchema(models.LtxSchema, dbname)
	if err == nil {
		// open the database
		var db *sqlx.DB
		db, err = sqlx.Open("sqlite3", dbname)
		if err != nil {
			logger.Println(err.Error())
			return errors.New(err.Error() + ": " + dbname)
		}
		idMap := make(map[string]bool)

		tx, err := db.Begin()
		check(err,logger)
		for ltx := range mappers.Lp2Lt(repos.Ares2LpFromGithub(urls, "ares", printProgress, logger)) {
			// report non-unique id
			if idMap[ltx.ID] {
				logger.Println(fmt.Sprintf("Duplicate ID: %s", ltx.ID))
			} else { // write sql for unique id
				idMap[ltx.ID] = true
				_, err = tx.Exec(fmt.Sprintf(models.ReadSQLInsert, ltx.ID, ltx.Library, ltx.Topic, ltx.Key, strings.ReplaceAll(ltx.Value, "'", "''"), strings.ReplaceAll(ltx.NNP, "'", "''"), strings.ReplaceAll(ltx.NWP, "'", "''"), strings.ReplaceAll(ltx.Comment, "'", "''"), ltx.Redirect, ltx.CreatedWhen, ltx.ModifiedWhen))
				check(err,logger)
			}
		}
		err = tx.Commit()
	}
	return err
}
// Reads lines from ares files in local directories,
// then writes the lines from each file into the database.
func AresDir2Sqlite(
	rootDir string,
	dbname string,
	printProgress bool,
	logger *log.Logger,
) error {
	var err error
	// delete existing database
	if FileExists(dbname) {
		os.Remove(dbname)
	}
	// set up the schema
	err = CreateSchema(models.LtxSchema, dbname)
	if err == nil {
		// open the database
		var db *sqlx.DB
		db, err = sqlx.Open("sqlite3", dbname)
		if err != nil {
			logger.Println(err.Error())
			return errors.New(err.Error() + ": " + dbname)
		}
		// pipeline: each line of each file of each repo is processed as follows...
		// Ares2LpFromLocalDir -> LineParts -> Lp2Lt -> Ltx, which is then written to sqlite
		for ltx := range mappers.Lp2Lt(repos.Ares2LpFromLocalDir(rootDir, "ares", printProgress, logger)) {
			_, err := db.Exec(models.LtxSQLInsert, ltx.ID, ltx.Library, ltx.Topic, ltx.Key, ltx.Value, ltx.NNP, ltx.NWP, ltx.Comment, ltx.Redirect, ltx.CreatedWhen, ltx.ModifiedWhen)
			if err != nil {
				if strings.HasPrefix(err.Error(), "UNIQUE") {
					logger.Printf("duplicate key: %s ", ltx.ID)
				} else {
					logger.Printf(err.Error())
				}
			}
		}
	}
	return err
}
// Check for an error.  If found, log it.
func check(err error, logger *log.Logger) {
	if err != nil {
		logger.Println(err.Error())
	}
}

