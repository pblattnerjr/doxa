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
// TODO: once AGES is using Doxa instead of ALWB, this command may be removed
package cmd

import (
	"fmt"
	"github.com/liturgiko/doxa/pkg/temp/atem2liml"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)


var atem2lmlCmd = &cobra.Command{
	Use:   "atem2lml",
	Short: "convert atem files to lml files",
	Long: `convert AGES atem template files to liturgical markup language (lml) files`,
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

		msg := fmt.Sprintf("converting atem to lml...\n")
		fmt.Println(msg)
		Logger.Println(msg)
		dirIn := "~/git/ages/atem/ages-alwb-templates/net.ages.liturgical.workbench.templates/a-templates"
		if err = atem2mt.Process(dirIn, Paths.TemplatesPath); err != nil {
			Logger.Println(err.Error())
		}
		Elapsed(start)
	},
}

func init() {
	rootCmd.AddCommand(atem2lmlCmd)
}

