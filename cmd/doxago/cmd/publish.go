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
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)


var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "publish the local website to one on the Internet",
	Long: `publish the local website to one on the Internet by adding, committing, and pushing the files to a github repo tied to a github site.
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

		msg := fmt.Sprintf("publish website is not implemented\n")
		fmt.Println(msg)
		Logger.Println(msg)
		Elapsed(start)
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)
}

