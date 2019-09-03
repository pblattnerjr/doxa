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
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strconv"
	"strings"
	"syscall"
	"text/template"
)

var whereClause string

const (
	CYPHER = "CYPHER$ "
	WHERE  = "WHERE$ "
)
type IdValue struct {
	ID string
	Val string
}
const tmplIdValue = `+------------------------------------+
ID : {{.ID}} 
Val: {{.Val}}
`
var t = template.Must(template.New("record").Parse(tmplIdValue))

// qneoCmd represents the qneo command
var qneoCmd = &cobra.Command{
	Use:   "qneo",
	Short: "Query neo4j database",
	Long:  `Query a neo4j database that contains liturgical texts, local or remote, using login credentials. Commands: .exit to quit, .cypher for ability to enter a full cypher command, .where to just enter the WHERE clause`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Host domain: ")
		var host string
		fmt.Scanln(&host)
		fmt.Print("Username: ")
		var username string
		fmt.Scanln(&username)
		fmt.Print("Password: ")
		bytepwd, _ := terminal.ReadPassword(int(syscall.Stdin))
		pwd := string(bytepwd)
		fmt.Println("")

		var (
			driver  neo4j.Driver
			session neo4j.Session
			result  neo4j.Result
			err     error
		)

		if driver, err = neo4j.NewDriver(fmt.Sprintf("bolt://%s:7687", host), neo4j.BasicAuth(username, pwd, "")); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		// handle driver lifetime based on your application lifetime requirements
		// driver's lifetime is usually bound by the application lifetime, which usually implies one driver instance per application
		defer driver.Close()

		if session, err = driver.Session(neo4j.AccessModeWrite); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		defer session.Close()

		_, err = session.Run("match (n) return n limit 1", nil)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		consoleReader := bufio.NewReader(os.Stdin)

		var prompt = "WHERE$ "

		for {
			fmt.Print(prompt)
			input, _ := consoleReader.ReadString('\n')
			input = input[:len(input)-1] // get rid of line feed
			switch strings.ToLower(input) {
			case ".exit":
				Doxa()
				os.Exit(0)
			case ".cypher":
				prompt = CYPHER
			case ".where":
				prompt = WHERE
			default: // process query
				q := formatQuery(prompt, input)
				result, err = session.Run(q, nil)
				if err != nil {
					fmt.Printf(err.Error())
				}
				for result.Next() {
					record := result.Record()
					if prompt == CYPHER {
						fmt.Println(fmt.Sprintf("%v", record.Values()))
						fmt.Println("+----------------------------------------+")
					} else {
						id, _ := record.Get("n.id")
						value, _ := record.Get("n.value")
						encode, _ := json.Marshal(value)
						a, _ := strconv.Unquote(string(encode))
						r := IdValue{fmt.Sprintf("%s",id), a}
						if err := t.Execute(os.Stdout, r); err != nil {err.Error()}
					}
				}
				fmt.Println(q)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(qneoCmd)
	qneoCmd.PersistentFlags().StringVar(&whereClause, "where", "", "")
}

func formatQuery(prompt, input string) string {
	var result string
	switch prompt {
	case CYPHER:
		result = input
	case WHERE:
		result = fmt.Sprintf("MATCH (n:Liturgical) WHERE %s return n.id, n.value order by n.id", input)
	}
	return result
}
