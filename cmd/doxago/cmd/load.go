// Copyright © 2019 The Orthodox Christian Mission Center (ocmc.org)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/**
Bugs
1. The total records are different between running load ares from dir vs from github by about 400
*/

package cmd

import (
	"bufio"
	"fmt"
	"github.com/liturgiko/doxa/pkg/config"
	"github.com/liturgiko/doxa/pkg/db/lsql"
	"github.com/liturgiko/doxa/pkg/utils/oslw"
	"github.com/liturgiko/doxa/pkg/utils/repos"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
)

var dbFilename = "liturgical.db"

// vars for 'load' flags
var (
	loadAres   bool
	loadJson   bool
	loadCvs    bool
	loadTexRes bool
	loadTexTmp bool
	ares       = "ares"
	cvs        = "cvs"
	texRes     = "OSLW Resources"
	texTmp     = "OSLW Templates"
	whatJson   = "json"
)

// vars for 'from' flags
var (
	fromDir    bool
	fromGithub bool
	fromGitlab bool
	dir        = "dir"
	directory = "directory"
	github     = "github"
	gitlab     = "gitlab"
)

// vars for 'to' flags
var (
	toSql bool
	toNeo bool
	sql = "sql"
	sqlLong = "sqlite"
	neo = "neo"
	neoLong = "neo4j"
)

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load a resource from a source to a destination",
	Long: `Load a resource (for example, an AGES ares file
from a source (e.g. a directory of git repositories) 
to a destination (e.g. a database). You can provide the 
information for the load in one of two ways.  
You can enter it as follows:
load ares from github to sql
or, you can just use the load command, and it will prompt
you to enter each required piece of information, displaying
the valid options at each point.
`,
	Run: func(cmd *cobra.Command, args []string) {

		start := time.Now()

		// set up the logger, which will be passed to the functions that do the processing
		LogFile, err := os.OpenFile(LogFilename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			panic(err)
		}
		defer LogFile.Close()

		Logger.SetOutput(LogFile)
		Logger.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)

		setParameters(args)

		// get the flags
		printProgress, _ := cmd.Flags().GetBool("verbose")
		useTestData, _ := cmd.Flags().GetBool("test")
		var theRepos []string
		flagRepo, _ := cmd.Flags().GetString("cloneUrl")
		if flagRepo != "" {
			theRepos = append(theRepos, flagRepo)
		} else {
			if useTestData {
				theRepos = viper.GetStringSlice("test.github.repos.ares")
			} else {
				theRepos = viper.GetStringSlice("github.repos.ares")
			}
		}
		aresPath := path.Join(Paths.ReposPath, "ares")

		msg := fmt.Sprintf("loading %s from %s to %s", loadWhat(), loadFrom(), loadTo())
		fmt.Println(msg)
		Logger.Println(msg)

		if toSql {
			dbFilename = filepath.Join(DOXAHOME, config.DataDir, config.SQL, dbFilename)
			fmt.Println(fmt.Sprintf("DB file is: %s", dbFilename))
		}

		// based on the parameters, determine which function to call
		switch {
		case loadAres:
			switch {
			case fromDir:
				msg = fmt.Sprintf("Reading from directory %s",aresPath)
				fmt.Println(msg)
				Logger.Println(msg)
				switch {
				case toNeo:
				case toSql:
					err := lsql.AresDir2Sqlite(aresPath, dbFilename, printProgress, &Logger)
					if err != nil {
						fmt.Println(err.Error())
					}
					fmt.Printf("\nFinished loading ares files into %s.", dbFilename)
					fmt.Printf("\nCheck %s to see if there were errors.\n", LogFilename)
				}
			case fromGithub:
				switch {
				case toNeo:
				case toSql:
					err := lsql.AresGithub2Sqlite(theRepos, dbFilename, printProgress, &Logger)
					if err != nil {
						fmt.Println(err.Error())
					}
					fmt.Printf("\nFinished loading ares files into %s.", dbFilename)
					fmt.Printf("\nCheck %s to see if there were errors.\n", LogFilename)
				default:
					fmt.Println("What to load into is not defined.  Exiting.")
					os.Exit(1)
				}
			case fromGitlab:
			default:
				fmt.Println("What to load from is not defined.  Exiting.")
				os.Exit(1)
			}
		case loadCvs:
		case loadJson:
		case loadTexRes:
			oslwPath := path.Join(Paths.ReposPath, "oslw")
			switch {
			case fromDir:
				msg = fmt.Sprintf("Reading from directory %s",oslwPath)
				fmt.Println(msg)
				Logger.Println(msg)
				switch {
				case toNeo:
				case toSql:
					err = oslw.Res2Sql(oslwPath, dbFilename, &Logger)
					if err != nil {
						fmt.Println(err.Error())
					}
					fmt.Printf("\nFinished loading OSLW resources into %s.", dbFilename)
					fmt.Printf("\nCheck %s to see if there were errors.\n", LogFilename)
				}
			case fromGithub:
				oslwCloneUrl := viper.GetString("github.repos.oslw")
				if len(oslwCloneUrl) == 0 {
					fmt.Println("Could not find OSLW clone url in .doxago.yaml")
					os.Exit(1)
				}
				_, err := repos.Clone(oslwPath, oslwCloneUrl, true)
				switch {
				case toSql:
					// TODO: can't call Reps2Sqlite for 2 reasons: 1) oslwCloneUrl is not an array; 2) that function expects ares.
					err := lsql.AresGithub2Sqlite(theRepos, dbFilename, printProgress, &Logger)
					if err != nil {
						fmt.Println(err.Error())
					}
					fmt.Printf("\nFinished loading ares files into %s.", dbFilename)
					fmt.Printf("\nCheck %s to see if there were errors.\n", LogFilename)
				default:
					fmt.Println("What to load into is not defined.  Exiting.")
					os.Exit(1)
				}
			default:
				fmt.Println("What to load from is not defined.  Exiting.")
				os.Exit(1)
			}
		default:
			fmt.Println("What to load is not defined.  Exiting.")
			os.Exit(1)
		}
		Elapsed(start)
	},
}

func init() {
	rootCmd.AddCommand(loadCmd)
}

func setParameters(args []string) {
	// see if we can get the information to load from the arguments
	whatOk, fromOk, toOk := setFromArgs(args)
	// prompt the user to enter arguments not provided, or that were invalid
	if ! whatOk {
		setWhat()
	}
	if ! fromOk {
		setFrom(loadWhat())
	}
	if ! toOk {
		setTo(loadWhat(), loadFrom())
	}
}
// Set the parameters using the command line arguments
// Each argument that is valid will return a true
// Any argument that is not valid will return a false
func setFromArgs(args []string) (bool, bool, bool) {
	var whatOk bool
	var fromOk bool
	var toOk bool
	if len(args) > 0 {
		whatOk = setWhatFromArg(args[0])
	}
	if len(args) > 2 {
		fromOk = setFromFromArg(args[2])
	}
	if len(args) > 4 {
		toOk = setToFromArg(args[4])
	}
	return whatOk, fromOk, toOk
}
func setWhatFromArg(arg string) bool {
	switch arg {
	case ares: {
		loadAres = true
		return true
	}
	case cvs: {
		loadCvs = true
		return true
	}
	case texRes: {
		loadTexRes = true
		return true
	}
	case whatJson: {
		loadJson = true
		return true
	}
	}
	return false
}
func setFromFromArg(arg string) bool {
	switch arg {
	case dir, directory:
		fromDir = true
		return true
	case github:
		fromGithub = true
		return true
	case gitlab:
		fromGitlab = true
		return true
	}
	return false
}
func setToFromArg(arg string) bool {
	switch arg {
	case neo, neoLong:
		toNeo = true
		return true
	case sql, sqlLong:
		toSql = true
		return true
	}
	return false
}
func setWhat() {
	reader := bufio.NewReader(os.Stdin)
l:
	for {
		fmt.Println("load what?")
		fmt.Println("a) ares")
		fmt.Println("b) cvs")
		fmt.Println("c) json")
		fmt.Println("d) oslw resources")
		fmt.Println("e) oslw templates")
		fmt.Println("q) quit the command")
		char, _, _ := reader.ReadRune()
		switch char {
		case 'a':
			loadAres = true
			break l
		case 'b':
			loadCvs = true
			break l
		case 'c':
			loadJson = true
			break l
		case 'd':
			loadTexRes = true
			break l
		case 'e':
			loadTexTmp = true
			break l
		case 'q':
			os.Exit(0)
		default:
			fmt.Println("Invalid selection...")
		}
	}
}
func setFrom(what string) {
	reader := bufio.NewReader(os.Stdin)
l:
	for {
		fmt.Printf("load %s from what?\n", what)
		fmt.Println("a) dir (directory on your computer)")
		fmt.Println("b) github")
		fmt.Println("c) gitlab")
		fmt.Println("q) quit the command")
		char, _, _ := reader.ReadRune()
		switch char {
		case 'a':
			fromDir = true
			break l
		case 'b':
			fromGithub = true
			break l
		case 'c':
			fromGitlab = true
			break l
		case 'q':
			os.Exit(0)
		default:
			fmt.Println("Invalid selection...")
		}
	}
}

func setTo(what string, from string) {
	reader := bufio.NewReader(os.Stdin)
l:
	for {
		fmt.Printf("load %s from %s into what?\n", what, from)
		fmt.Println("a) neo (neo4j database)")
		fmt.Println("b) sql (sqlite database)")
		fmt.Println("q) quit the command")
		char, _, _ := reader.ReadRune()
		switch char {
		case 'a':
			toNeo = true
			break l
		case 'b':
			toSql = true
			break l
		case 'q':
			os.Exit(0)
		default:
			fmt.Println("Invalid selection...")
		}
	}
}
func loadWhat() string {
	switch {
	case loadAres:
		return ares
	case loadCvs:
		return cvs
	case loadJson:
		return whatJson
	case loadTexRes:
		return texRes
	case loadTexTmp:
		return texTmp
	default:
		return "unknown"
	}
}
func loadFrom() string {
	switch {
	case fromDir:
		return "dir"
	case fromGithub:
		return "github"
	case fromGitlab:
		return "gitlab"
	default:
		return "unknown"
	}
}
func loadTo() string {
	switch {
	case toSql:
		return "sqlite database"
	case toNeo:
		return "neo4j database"
	default:
		return "unknown"
	}
}
