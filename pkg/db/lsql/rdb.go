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
	"github.com/liturgiko/doxa/pkg/utils/ares"
	"github.com/liturgiko/doxa/pkg/utils/ltstring"
	"github.com/liturgiko/doxa/pkg/utils/repos"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"
	"log"
	"net/http"
	"os"
	//	"path/filepath"
	"strings"
	//	"text/template"
)


// Schema to create Ltext
var SchemaLtext = `CREATE TABLE ltext (
    id       TEXT PRIMARY KEY,
    topic    TEXT,
    key      TEXT,
    value    TEXT,
    nnp      TEXT,
    nwp      TEXT,
    comment  TEXT,
    redirect TEXT);`

// SQL to insert Ltext into database
var InsertLtext = `INSERT INTO ltext (id, topic, key, value, nnp, nwp, comment, redirect) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// For each repo, clones in memory, reads lines in files and saves to a sqlite database.
// If the database file exists, it will first be deleted.
func ReposToSqlite(
	repos []string,
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
	err = CreateSchema(SchemaLtext, dbname)
	if err == nil {
		for _, repo := range repos {
			err := RepoToSqlite(repo, dbname, printProgress, logger)
			if err != nil {
				logger.Println(err.Error())
			}
		}
	}
	return err
}

// Create a schema in the specified database
func CreateSchema(schema string, dbname string) error {
	db, err := sqlx.Connect("sqlite3", dbname)
	_, err = db.Exec(schema)
	return err
}

// Loads the repo into memory from Github, then writes the lines from each file into the database.
func RepoToSqlite(
	repo string,
	dbname string,
	printProgress bool,
	logger *log.Logger,
) error {
	var err error

	// open the database
	var db *sqlx.DB
	db, err = sqlx.Open("sqlite3", dbname)
	if err != nil {
		logger.Println(err.Error())
		return errors.New(err.Error() + ": " + dbname)
	}

	_, err = http.Get(repo)
	if err != nil {
		fmt.Println(fmt.Sprintf("No Internet connection, or bad url: %s", repo))
		return errors.New(err.Error() + ": " + repo)
	}
	// clone the repo into memory
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: repo,
	})
	if err != nil {
		// try again because it might have been caused by a netword delay (http request timeout)
		// TODO: find a way to specify a timeout and increase it.
		r, err = git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
			URL: repo,
		})
		if err != nil {
			return errors.New(err.Error() + ": " + repo)
		}
	}

	ref, err := r.Head()
	if err != nil {
		return errors.New(err.Error() + ": " + repo)
	}

	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		return errors.New(err.Error() + ": " + repo)
	}

	// ... retrieve the tree from the commit
	tree, err := commit.Tree()
	if err != nil {
		return errors.New(err.Error() + ": " + repo)
	}

	// ... get the files iterator and print the file
	err = tree.Files().ForEach(func(f *object.File) error {
		if strings.HasSuffix(f.Name, "ares") {
			var lineCnt int = 0
			var fileNameParts ares.FilenameParts
			fileNameParts, err := ares.ParseAresFileName(f.Name)
			if printProgress {
				fmt.Printf("\r%-80s", "")
				fmt.Print(fmt.Sprintf("\rProcessing %s_%s_%s_%s.ares", fileNameParts.Topic, fileNameParts.Language, fileNameParts.Country, fileNameParts.Realm))
			}

			if err == nil {
				l, _ := f.Lines()

				inCommentBlock := false // To handle multi-line /*.....*/

				for _, line := range l {
					lineCnt = lineCnt + 1
					line = strings.TrimSpace(line)
					if strings.HasPrefix(line, "/*") {
						inCommentBlock = true
						continue
					}
					if strings.HasPrefix(line, "*/") || strings.HasSuffix(line, "*/") {
						inCommentBlock = false
						continue
					}
					if !inCommentBlock {
						lineParts, err := ares.ParseLine(fileNameParts, line)
						if err != nil {
							logger.Printf("%s: %s: line %d", err.Error(), f.Name, lineCnt)
							continue
						}
						if lineParts.IsBlank || lineParts.IsCommentedOut {
							// ignore
						} else {
							var ltext models.Ltext
							ltext.ID = ltstring.ToId(
								fileNameParts.Language,
								fileNameParts.Country,
								fileNameParts.Realm,
								fileNameParts.Topic,
								lineParts.Key,
							)
							if lineParts.HasValue {
								ltext.Value = lineParts.Value
								ltext.NWP = ltstring.ToNwp(lineParts.Value)
								ltext.NNP = ltstring.ToNnp(ltext.NWP)
							}
							if lineParts.HasComment {
								ltext.Comment = lineParts.Comment
							}
							if lineParts.IsRedirect {
								ltext.Redirect = lineParts.Redirect
							}
							_, err := db.Exec(InsertLtext, ltext.ID, ltext.Topic, ltext.Key, ltext.Value, ltext.NNP, ltext.NWP, ltext.Comment, ltext.Redirect)
							if err != nil {
								if strings.HasPrefix(err.Error(), "UNIQUE") {
									logger.Printf("duplicate key: %s: line %d", f.Name, lineCnt)
								} else {
									logger.Printf("db error %s: %s: line %d", err.Error(), f.Name, lineCnt)
								}
							}
						}
					}
				}
			} else {
				logger.Printf("%s, file: %s\n", err.Error(), f.Name)
			}
		}
		return err
	})
	return err
}
// Loads the repo into memory from Github, then writes the lines from each file into the database.
func Repos2Sqlite(
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
	err = CreateSchema(SchemaLtext, dbname)
	if err == nil {
		// open the database
		var db *sqlx.DB
		db, err = sqlx.Open("sqlite3", dbname)
		if err != nil {
			logger.Println(err.Error())
			return errors.New(err.Error() + ": " + dbname)
		}
		//		var Data []*models.Ltext
		// pipeline: each line of each file of each repo is processed as follows...
		// Ares2LpFromGithub -> LineParts -> Lp2Lt -> Ltext, which is then written to sqlite
		//		db.Ping()
		for ltext := range mappers.Lp2Lt(repos.Ares2LpFromGithub(urls, "ares", printProgress, logger)) {
			//ltext.Value = strings.ReplaceAll(ltext.Value, "'","''")
			//ltext.NWP = strings.ReplaceAll(ltext.NWP, "'","''")
			//Data = append(Data, ltext)
			_, err := db.Exec(InsertLtext, ltext.ID, ltext.Topic, ltext.Key, ltext.Value, ltext.NNP, ltext.NWP, ltext.Comment, ltext.Redirect)
			if err != nil {
				if strings.HasPrefix(err.Error(), "UNIQUE") {
					logger.Printf("duplicate key: %s ", ltext.ID)
				} else {
					logger.Printf(err.Error())
				}
			}
		}
		// The idea below is to write a sql file, then load it into
		// an empty sqlite db.  If it is properly formatted, and does
		// not have any errors (e.g. duplicate ids), it loads in about 3 seconds.
		// But, at the moment, there are bad lines being written that prevent the load
		// from finishing.
		//t := template.Must(template.New("sql").Parse(sqlFileTmpl))
		//f, err := os.Create("/Users/mac002/doxago/data/sql/out.sql")
		//if err != nil {
		//	log.Println("create file: ", err)
		//}
		//err = t.Execute(f,Data)
	}
	return err
}
// Reads lines from ares files in local directories,
// then writes the lines from each file into the database.
func LocalAres2Sqlite(
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
	err = CreateSchema(SchemaLtext, dbname)
	if err == nil {
		// open the database
		var db *sqlx.DB
		db, err = sqlx.Open("sqlite3", dbname)
		if err != nil {
			logger.Println(err.Error())
			return errors.New(err.Error() + ": " + dbname)
		}
		// pipeline: each line of each file of each repo is processed as follows...
		// Ares2LpFromLocalDir -> LineParts -> Lp2Lt -> Ltext, which is then written to sqlite
		for ltext := range mappers.Lp2Lt(repos.Ares2LpFromLocalDir(rootDir, "ares", printProgress, logger)) {
			_, err := db.Exec(InsertLtext, ltext.ID, ltext.Topic, ltext.Key, ltext.Value, ltext.NNP, ltext.NWP, ltext.Comment, ltext.Redirect)
			if err != nil {
				if strings.HasPrefix(err.Error(), "UNIQUE") {
					logger.Printf("duplicate key: %s ", ltext.ID)
				} else {
					logger.Printf(err.Error())
				}
			}
		}
	}
	return err
}
const sqlFileTmpl = `PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE ltext (
    id       TEXT PRIMARY KEY,
    topic    TEXT,
    key      TEXT,
    value    TEXT,
    nnp      TEXT,
    nwp      TEXT,
    comment  TEXT,
    redirect TEXT);
{{ range .}}INSERT INTO ltext VALUES('{{.ID}}','{{.Topic}}','{{.Key}}','{{.Value}}','{{.NNP}}','{{.NWP}}','{{.Comment}}','{{.Redirect}}');
{{end}}COMMIT;
`