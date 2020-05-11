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
	"github.com/liturgiko/doxa/pkg/utils/repos"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"time"
)


var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "clone the github repos listed in the config file",
	Long: `clone the github repositories listed in the config file
that are identified by the key github.repos, to the ares.dir (directory)
value setRecord in the config file. Be aware that if the directory exists,
it will first be deleted.
`,
	Run: func(cmd *cobra.Command, args []string) {

		start := time.Now()

		// setRecord popPath the logger, which will be passed to the functions that do the processing
		LogFile, err := os.OpenFile(LogFilename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			panic(err)
		}
		defer LogFile.Close()

		Logger.SetOutput(LogFile)
		Logger.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)

		var repoPath string
		// get the flags
		repoType, _ := cmd.Flags().GetString("repo")
		useTestData, _ := cmd.Flags().GetBool("test")
		var theRepos []string
		flagRepo, _ := cmd.Flags().GetString("cloneUrl")
		if flagRepo != "" {
			theRepos = append(theRepos, flagRepo)
		} else {
			if useTestData {
					theRepos = viper.GetStringSlice("test.github.repos." + repoType)
			} else {
				theRepos = viper.GetStringSlice("github.repos." + repoType)
			}
		}
		repoPath = filepath.Join(REPOPATH, repoType)

		if len(theRepos) == 0 {
			fmt.Printf("Repository list %s not found in config file\n", repoType)
		}

		msg := fmt.Sprintf("cloning into %s", repoPath)
		fmt.Println(msg)
		Logger.Println(msg)

		if err = repos.CloneConcurrently(repoPath, theRepos); err != nil {
			Logger.Println(err.Error())
		}
		Elapsed(start)
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
	cloneCmd.Flags().StringP("repo", "r", "site", "type of repositories to process, e.g. ares, site. For example, doxago clone -r ares")
}

