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
//
// When running doxa in the cloud via docker-compose, set environmental
// variables as follows:
// cloud true
// apiport {port number}, e.g. 8080
// appport {port number}, e.g. 8090
//
// Then, set up nginx (or whatever) to proxy the server running on the appport

package cmd

import (
	"fmt"
	"github.com/liturgiko/doxa/pkg/config"
	"github.com/liturgiko/doxa/pkg/db/lsql"
	"github.com/liturgiko/doxa/pkg/server/app"
	"github.com/liturgiko/doxa/pkg/utils/ltfile"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"strings"
	"time"
)

var dbFilename = "liturgical.db"
var DOXAHOME string
var DOXAPORT = ""
var APIPORT = ""
var GLORY string
var REPOPATH string

var Paths *config.Paths

var cfgFile = ".doxa"
var Logger log.Logger
var LogFile *os.File
var LogFilename string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "doxa",
	Short: "Web browser app for creating and managing Eastern Orthodox Christian liturgical texts",
	Long: `doxa provides a web browser based user interface you can use to create 
and manage liturgical texts in multiple languages.
++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
doxa is a web application written in the Go programming language.

This application is a tool to create and manage multilingual 
liturgical texts used by the Eastern Orthodox Christian Church.  
It was created by staff and volunteers of the 
Orthodox Christian Mission Center (OCMC), and is provided as 
a free tool to the pan-Orthodox community.  

doxa always gives glory to God!

++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	var cloud = strings.ToLower(os.Getenv("cloud"))
	inCloud := cloud == "true"
	initConfig()
	DOXAPORT = os.Getenv("appport")
	if len(DOXAPORT) == 0 {
		DOXAPORT = viper.GetString("port.http.doxa")
	}
	if len(DOXAPORT) == 0 {
		DOXAPORT = "8080"
	}
	APIPORT = os.Getenv("apiport")
	if len(APIPORT) == 0 {
		APIPORT = "8090"
	}
	app.Serve(Paths.DbPath,"8080", "8090", inCloud)
}

// initConfig reads in config file and ENV variables if setRecord.
func initConfig() {
	homeDir := homeDir()
	DOXAHOME = path.Join(homeDir,"doxa")

	// if the DOXA home does not exist, initialize it
	if _, err := os.Stat(DOXAHOME); os.IsNotExist(err) {
		initializeDoxaHome()
	} else { // if it does exist, but the config is missing, initialize again
		if _, err := os.Stat(path.Join(DOXAHOME,"doxago.yaml")); os.IsNotExist(err) {
			initializeDoxaHome()
		}
	}

	DOXAPORT = viper.GetString("port.http.doxa")

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found
			fmt.Println("Could not find config file.", viper.ConfigFileUsed())
		} else {
			// Config file was found but another error was produced
			fmt.Println("Found config file, but had error", err.Error())
		}
	} else {
		os.Mkdir(DOXAHOME, os.ModePerm)
		LogFilename = filepath.Join(DOXAHOME, config.LogDir, "doxago.log")
	}
	gloryLangs := viper.GetStringSlice("glory.languages")
	gloryIndex := viper.GetInt("glory")
	if gloryIndex > len(gloryLangs) {
		fmt.Println("The glory index is greater than the number of languages.  The index starts with 0.")
	} else {
		GLORY = gloryLangs[gloryIndex]
	}
	REPOPATH = filepath.Join(DOXAHOME, "repos")
	ltfile.CreateDir(REPOPATH)
	Paths = config.NewPaths(DOXAHOME, viper.GetString("github.repos.site"))
	db :=	filepath.Join(DOXAHOME, config.DataDir, config.SQL, dbFilename)

	if _, err := os.Stat(db); os.IsNotExist(err) {
		initializeDb(db)
	}
}

func Elapsed(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Took %s\n", elapsed)
	Doxa()
}
func Doxa() {
	fmt.Println(GLORY)
}

// homeDir returns the home directory of the current user
func homeDir() string {
	usr, err := user.Current()
	if err != nil {
		Logger.Println(err.Error())
	}
	return usr.HomeDir
}

func initializeDoxaHome() {
	r, err := config.Initdoxa(DOXAHOME)

	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	} else {
		fmt.Println(r)
	}
	// Search config in home directory with name ".doxa" (without extension).
	viper.AddConfigPath(DOXAHOME)
	viper.SetConfigName(".doxago")
}

func initializeDb(db string) {
	// load the database
	start := time.Now()

	// setRecord popPath the logger, which will be passed to the functions that do the processing
	LogFile, err := os.OpenFile(LogFilename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer LogFile.Close()

	Logger.SetOutput(LogFile)
	Logger.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)

	repos := viper.GetStringSlice("github.repos.ares")
	fmt.Println("loading ares files from Github to sql database")
	fmt.Println(fmt.Sprintf("DB file is: %s", db))
	err = lsql.AresGithub2Sqlite(repos, db, true, &Logger)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("\nFinished loading ares files into %s.", dbFilename)
	fmt.Printf("\nCheck %s to see if there were errors.\n", LogFilename)
	Elapsed(start)

}