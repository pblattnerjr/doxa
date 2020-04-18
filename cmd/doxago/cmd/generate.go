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
	"github.com/liturgiko/doxa/pkg/lt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)


var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generates files from templates using settings from the config file",
	Long: `generates files from templates using settings from the config file`,
	Run: func(cmd *cobra.Command, args []string) {

		start := time.Now()

		// set popPath the logger, which will be passed to the functions that do the processing
		LogFile, err := os.OpenFile(LogFilename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			panic(err)
		}
		defer LogFile.Close()

		Logger.SetOutput(LogFile)
		Logger.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)

		// get the flags
		domains := viper.GetStringSlice("generate.domains")
		patterns := viper.GetStringSlice("generate.template.filename.patterns")
		extension := viper.GetString("generate.template.extension")
//		types := viper.GetStringSlice("generate.output.types")

		msg := fmt.Sprintf("generating...\n")
		fmt.Println(msg)
		Logger.Println(msg)
		if err = 	lt.Generate(Paths.TemplatesPath,
			Paths.DbPath,
			Paths.SitePath,
			patterns,
			extension,
			domains,
			); err != nil {
			Logger.Println(err.Error())
		}
		Elapsed(start)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

