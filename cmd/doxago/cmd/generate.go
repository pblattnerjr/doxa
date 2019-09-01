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
	"time"
)


var generateCmd = &cobra.Command{
	Use:   "clone",
	Short: "clones the github repos listed in the config file",
	Long: `clones the github repositories listed in the config file
that are identified by the key github.repos, to the ares.dir (directory)
value set in the config file. Be aware that if the directory exists,
it will first be deleted.
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

		// get the flags
		useTestData, _ := cmd.Flags().GetBool("test")
		var repos []string
		flagRepo, _ := cmd.Flags().GetString("cloneUrl")
		if flagRepo != "" {
			repos = append(repos, flagRepo)
		} else {
			if useTestData {
				repos = viper.GetStringSlice("test.github.repos")
			} else {
				repos = viper.GetStringSlice("github.repos")
			}
		}
		aresPath := viper.GetString("ares.dir")

		msg := fmt.Sprintf("cloning into %s", aresPath)
		fmt.Println(msg)
		Logger.Println(msg)

		if err = repos2.CloneConcurrently(aresPath, repos); err != nil {
			Logger.Println(err.Error())
		}

		elapsed := time.Since(start)
		log.Printf("Took %s", elapsed)

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

