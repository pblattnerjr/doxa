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

package cmd

import (
	"fmt"
	"github.com/liturgiko/doxa/pkg/config"
	"github.com/liturgiko/doxa/pkg/utils/ltfile"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

var DOXAHOME string
var LIMLHOME string
var DOXAPORT string
var GLORY string
var REPOPATH string

var Paths *config.Paths

var cfgFile = ".doxa"
var Logger log.Logger
var LogFile *os.File
var LogFilename string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "doxago",
	Short: "Tools for creating and managing Eastern Orthodox Christian liturgical texts",
	Long: `doxago provides commands you can use to create 
and manage liturgical texts in multiple languages.
++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
doxago is a CLI library written in the Go programming language.

This application is a tool to create and manage multilingual 
liturgical texts used by the Eastern Orthodox Christian Church.  
This tool was created by staff and volunteers of the 
Orthodox Christian Mission Center (OCMC).

It is provided as a free tool to the pan-Orthodox community.  

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
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.doxago.yaml)")
	rootCmd.PersistentFlags().StringVar(&LogFilename, "logfile", "doxago.log", "name of log file")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "display information while the command is running")
	rootCmd.PersistentFlags().BoolP("test", "t", false, "when set will use test data identified in the config file")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	DOXAHOME = os.Getenv("DOXAHOME")
	LIMLHOME = os.Getenv("LIMLHOME")
	DOXAPORT = os.Getenv("DOXAPORT")
	checkVar(DOXAHOME)
	checkVar(DOXAPORT)
	checkVar(LIMLHOME)

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
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
}
func checkVar(v string) {
	if v == "" {
		fmt.Println(tmpl)
		os.Exit(1)
	}
}
const tmpl = `
doxa cannot run unless you first set 
the environmental variables shown below.
How you do this depends on your operating system.

On a Mac, edit .bash_profile in your home directory and
add the following:

export DOXAPORT=8080
export DOXAHOME=$HOME/doxa
export PATH=$PATH:$DOXAHOME

On Windows, TBD...`
