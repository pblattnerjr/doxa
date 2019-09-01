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
	"github.com/liturgiko/doxa/pkg/css"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var (
	all bool
	initializeConfig bool
	initCss bool
	db bool
	templates bool
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize doxago",
	Long:  `initialize the liturgical text processor (doxago) folder and files it uses`,
	Run: func(cmd *cobra.Command, args []string) {
		if ! (all || initCss || db || templates) {
			cmd.Help()
			os.Exit(0)
		}
		if all {
			fmt.Println("Initializing everything")
			initCss, db, templates = true, true, true
		}
		if initCss {
			fmt.Println("Initializing css")
			css.WriteCss(filepath.Join(DOXAHome, "http","static"),"doxago.css")
		}
		if db {
			fmt.Println("Initializing db")
		}
		if templates {
			fmt.Println("Initializing templates")
		}
		fmt.Println("Finished...")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.PersistentFlags().BoolVarP(&all,"all","a",false,"initialize all assets")
	initCmd.PersistentFlags().BoolVarP(&initCss,"css","c",false,"initialize css")
	initCmd.PersistentFlags().BoolVarP(&initializeConfig,"config","f",false,"initialize css")
	initCmd.PersistentFlags().BoolVarP(&db,"db","d",false,"initialize database")
	initCmd.PersistentFlags().BoolVarP(&templates,"templates","",false,"initialize templates")
}

