// Copyright © 2020 The Orthodox Christian Mission Center (ocmc.org)

package cmd

import (
	"fmt"
	prompt "github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "provides a shell for accessing the liturgical database",
	Long:  `provides a shell for accessing the liturgical database`,
	Run: func(cmd *cobra.Command, args []string) {
		setSuggestions()
		p := prompt.New(
			executor,
			completer,
			prompt.OptionPrefix(thePath +"> "),
			prompt.OptionLivePrefix(livePrefix),
			prompt.OptionTitle("doxa-shell"),
		)
		p.Run()
	},
}
func init() {
	rootCmd.AddCommand(shellCmd)
}
var thePath = ""
const pathDelimiter = "/"
var suggestions []prompt.Suggest

func livePrefix() (string, bool) {
	if thePath == "" {
		return "", false
	}
	return thePath + "> ", true
}

func executor(in string) {
	in = strings.TrimSpace(in)

	var method string
	blocks := strings.Split(in, " ")
	switch blocks[0] {
	case "exit":
		fmt.Println("Bye!")
		os.Exit(0)
	case "cd":
		// 'cd' to home with empty path not yet supported.
		if len(blocks) < 2 {
			thePath = ""
		} else {
			if strings.HasPrefix(blocks[1],"..") {
				thePath = up(blocks[1])
				return
			}
			switch blocks[1] {
			case "~": thePath = ""
			default:
				if len(thePath) > 0 {
					thePath = thePath + pathDelimiter + blocks[1]
				} else {
					thePath = blocks[1]
				}
			}
		}
		return
	case "get", "delete":
		method = strings.ToUpper(blocks[0])
	case "post", "put", "patch":
		if len(blocks) < 2 {
			fmt.Println("please set request body.")
			return
		}
		method = strings.ToUpper(blocks[0])
	}
	if method != "" {
		return
	}

	if h := strings.Split(in, ":"); len(h) == 2 {
	} else {
		fmt.Println("Sorry, I don't understand.")
	}
}

func completer(in prompt.Document) []prompt.Suggest {
	w := in.GetWordBeforeCursor()
	if w == "" {
		return []prompt.Suggest{}
	}
	return prompt.FilterHasPrefix(suggestions, w, true)
}
func up(levels string) string {
	l := strings.Split(levels, pathDelimiter)
	s := strings.Split(thePath, pathDelimiter)
	if len(s) <= len(l) {
		return ""
	}
	var b strings.Builder
	nl := len(s) - len(l)
	for i := 0; i < nl; i++ {
		b.WriteString(s[i])
		if nl-i > 1 {
			b.WriteString(pathDelimiter)
		}
	}
	return b.String()
}
func setSuggestions() {
	suggestions = []prompt.Suggest{
		// Command
		{"cd", "CHANGE path"},
		{"cp", "COPY contents"},
		{"exact", "EXACT on, EXACT off. If on, find will match case and accents."},
		{"exit", "Exit Doxa Shell"},
		{"find", "FIND records with specified value"},
		{"get", "GET value for specified id, e.g. get en_us_dedes~actors~Priest"},
		{"ls", "LIST contents"},
		{"mv", "MOVE (rename) matching id to new id"},
		{"rm", "REMOVE for matching id"},
		{"set", "SET value for specified id, e.g. set en_us_dedes~actors~Priest=\"Priest\""},
	}
	appendLibraries()
	appendTopics()
}
func appendLibraries() {
	library := prompt.Suggest{Text: "en_us_dedes", Description: "English, USA, Fr. Seraphim Dedes"}
	suggestions = append(suggestions, library)
	library = prompt.Suggest{Text: "gr_gr_cog", Description: "Greek, Greece, Common Orthodox Greek"}
	suggestions = append(suggestions, library)
}
func appendTopics() {
	topic := prompt.Suggest{Text: "actors", Description: "Types of people with roles in liturgical service"}
	suggestions = append(suggestions, topic)
	topic = prompt.Suggest{Text: "prayers", Description: "Often used prayers"}
	suggestions = append(suggestions, topic)
}



