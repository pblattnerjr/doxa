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
	repos2 "github.com/liturgiko/doxa/pkg/utils/repos"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"time"
)


var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "push the github repos listed in the config file",
	Long: `push the github repositories listed in the config file
that are identified by the key github.repos.  It will first add *, 
then commit, using the date and time as the commit message.
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
		var repos []string
		flagRepo, _ := cmd.Flags().GetString("cloneUrl")
		if flagRepo != "" {
			repos = append(repos, flagRepo)
		} else {
			if useTestData {
					repos = viper.GetStringSlice("test.github.repos." + repoType)
			} else {
				repos = viper.GetStringSlice("github.repos." + repoType)
			}
		}
		repoPath = filepath.Join(REPOPATH, repoType)

		if len(repos) == 0 {
			fmt.Printf("Repository list %s not found in config file\n", repoType)
		}

		msg := fmt.Sprintf("cloning into %s", repoPath)
		fmt.Println(msg)
		Logger.Println(msg)

		if err = repos2.CloneConcurrently(repoPath, repos); err != nil {
			Logger.Println(err.Error())
		}
		Elapsed(start)
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
	pushCmd.Flags().StringP("repo", "r", "site", "type of repositories to process, e.g. ares, site.")
}

