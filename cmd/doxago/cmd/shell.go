// Copyright © 2020 The Orthodox Christian Mission Center (ocmc.org)

package cmd

import (
	SQL "database/sql"
	"fmt"
	prompt "github.com/c-bata/go-prompt"
	"github.com/liturgiko/doxa/pkg/db/ltx2sql"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"log"
	"os"
	"strconv"
	"strings"
)

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "provides a shell for accessing the liturgical database",
	Long:  `provides a shell for accessing the liturgical database`,
	Run: func(cmd *cobra.Command, args []string) {
		// open the database
		db, err := SQL.Open("sqlite3", Paths.DbPath)
		if err != nil {
			log.Println(err.Error())
		}
		defer db.Close()
		mapper.DB = db

		setSuggestions()
		p := prompt.New(
			executor,
			completer,
			prompt.OptionPrefix(context.path+"> "),
			prompt.OptionLivePrefix(livePrefix),
			prompt.OptionTitle("doxa-shell"),
		)
		p.Run()
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)
}

var mapper = ltx2sql.LtxMapper{}
var suggestions []prompt.Suggest
var pathDelimiter = "/" // strconv.QuoteRune(os.PathSeparator)
type Context struct {
	library string
	topic   string
	key     string
	path    string
}

var context Context
var previousContext Context

func (c *Context) Depth() int {
	if context.path == "" {
		return 0
	} else {
		return len(strings.Split(context.path,pathDelimiter))
	}
}
// Based on the depth of the path, provides a LIKE clause for sql select where like
func (c *Context) Like() string {
	like := ""
	switch c.Depth() {
	case 1: like = c.library + "~%"
	case 2: like = c.library + "~" + c.topic + "~%"
	case 3: like = c.library + "~" + c.topic + "~"  + c.key
	}
	return like
}
func (c *Context) Path() string {
	var b strings.Builder
	if len(context.key) > 0 {
		b.WriteString(context.library)
		b.WriteString(pathDelimiter)
		b.WriteString(context.topic)
		b.WriteString(pathDelimiter)
		b.WriteString(context.key)
		return b.String()
	}
	if len(context.topic) > 0 {
		b.WriteString(context.library)
		b.WriteString(pathDelimiter)
		b.WriteString(context.topic)
		return b.String()
	}
	return context.library
}

func livePrefix() (string, bool) {
	if context.path == "" {
		return "", false
	}
	return context.path + "> ", true
}

// executes the commands issued by the user
func executor(in string) {
	in = strings.TrimSpace(in)
	var method string
	blocks := strings.Split(in, " ")
	switch blocks[0] {
	case "exit":
		fmt.Println("Bye!")
		os.Exit(0)
	case "cd":
		copyContext()
		if len(blocks) < 2 {
			context.path = ""
		} else if blocks[1] == "~" {
			context.path = resetPath()
		} else {
			up, dirs := cdParts(blocks[1])
			if len(up) > 0 {
				if up == "~" {
					context.path = resetPath()
				} else {
					context.path = popPath(up)
				}
			}
			if len(dirs) > 0 {
				s := strings.Split(dirs, pathDelimiter)
				depth := context.Depth()
				if depth + len(s) > 3 {
					fmt.Println("You can't go deeper than 3 levels (library/topic/key).")
				} else {
					for _, segment := range s {
						context.path = pushPath(segment)
					}
				}
			}
			if context.Depth() == 3 {
				showValue()
			} else {
				reportCount()
			}
		}
		setSuggestions()
		return
	case "ls":
		method = strings.ToUpper(blocks[0])
		switch context.Depth() {
		case 0: { // list all libraries in database
			libraries, err := mapper.Libraries()
			if err != nil {
				Logger.Print(err)
			}
			for _, library := range libraries {
				fmt.Println(library)
			}
		}
		case 1: { // list topics for current library
			var like = context.library
			if len(blocks) > 1 {
				like = like + "~" + blocks[1]
			}
			topics, err := mapper.Topics(like)
			if err != nil {
				Logger.Print(err)
			}
			for _, topic := range topics {
				fmt.Println(topic)
			}
		}
		case 2: { // list keys for current library/topic
			var like = context.library + "~" + context.topic
			if len(blocks) > 1 {
				like = like + "~" + blocks[1]
			}
			topics, err := mapper.Keys(like)
			if err != nil {
				Logger.Print(err)
			}
			for _, topic := range topics {
				fmt.Println(topic)
			}
		}
		case 3: { // Display the value of the key
			showValue()
		}
		}
	case "set":
		// TODO: when ltx struct is modified to have properties for create/modify date and by whom, need to set these here.
		method = "set"
		rec, err := mapper.ReadByLTK(context.library, context.topic, context.key)
		if err != nil {
			fmt.Println(err)
		} else {
			value := ""
			if len(blocks) > 1 { // if len == 0, user wants to set value to blank
				// escape any embedded quotes
				value = strconv.Quote(strings.Join(blocks[1:len(blocks)], " "))
				// a side effect of strconv.Quote is it wraps the entire string in quotes,
				// which we do not want.  So, remove them.
				value = value[1 : len(value)-1]
			}
			rec.SetValue(value)
			mapper.Merge(rec)
			fmt.Println(rec.Value)
		}
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
func popPath(levels string) string {
	l := len(strings.Split(levels, pathDelimiter))
	s := len(strings.Split(context.path, pathDelimiter))
	r := s - l
	if r < 1 {
		context.library = ""
		context.topic = ""
		context.key = ""
	} else if r == 1 {
		context.topic = ""
		context.key = ""
	} else if r == 2 {
		context.key = ""
	} else {
		fmt.Println("Oops!")
	}
	return context.Path()
}
func pushPath(segment string) string {
	if len(context.library) < 1 {
		context.library = segment
	} else if len(context.topic) < 1 {
		context.topic = segment
	} else {
		context.key = segment
	}
	return context.Path()
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
	switch context.Depth() {
	case 0: appendLibraries()
	case 1: appendTopics()
	default: appendKeys()
	}
}
// get list of libraries from the database and append to suggestions
func appendLibraries() {
	libraries, err := mapper.Libraries()
	if err != nil {
		fmt.Println(err)
	}
	for _, library := range libraries {
		prompt := prompt.Suggest{Text: library, Description: ""}
		suggestions = append(suggestions, prompt)
	}
}
// get list of topics from the database and append to suggestions
func appendTopics() {
	items, err := mapper.Distinct("topic", context.Like())
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range items {
		prompt := prompt.Suggest{Text: item, Description: ""}
		suggestions = append(suggestions, prompt)
	}
}
// get list of topics from the database and append to suggestions
func appendKeys() {
		items, err := mapper.Distinct("key", context.Like())
		if err != nil {
			fmt.Println(err)
		}
		for _, item := range items {
			prompt := prompt.Suggest{Text: item, Description: ""}
			suggestions = append(suggestions, prompt)
		}
}
func resetPath() string {
	context.library = ""
	context.topic = ""
	context.key = ""
	return context.Path()
}
func cdParts(p string) (string, string) {
	var cdU = strings.Builder{}
	var cdD = strings.Builder{}
	parts := strings.Split(p, "/")
	for _, part := range parts {
		if len(part) == 0 {
			continue
		}
		if strings.HasPrefix(part, "..") || part == "~" {
			cdU.WriteString(part)
			cdU.WriteString("/")
		} else {
			cdD.WriteString(part)
			cdD.WriteString("/")
		}
	}
	u := cdU.String()
	if len(u) > 0 {
		u = u[:len(u)-1]
	}
	d := cdD.String()
	if len(d) > 0 {
		d = d[:len(d)-1]
	}
	return u, d
}
func reportCount() {
	if len(context.path) > 0 {
		p := message.NewPrinter(language.English)
		c, _ := mapper.CountKeys(context.library, context.topic, context.key)
		if c == 0 {
			fmt.Printf("Path %s does not exist\n", context.path)
			revertContext()
		} else {
			if context.Depth() == 1 {
				t, _ := mapper.CountTopics(context.library)
				p.Printf("%d topics %d keys\n",t, c)
			} else {
				p.Printf("%d keys\n",c)
			}
		}
	}
}
func copyContext() {
	previousContext.library = context.library
	previousContext.topic = context.topic
	previousContext.key = context.key
	previousContext.path = context.path
}
func revertContext() {
	context.library = previousContext.library
	context.topic = previousContext.topic
	context.key = previousContext.key
	context.path = previousContext.path
}
func showValue() {
	if context.Depth() == 3 {
		rec, err := mapper.ReadByLTK(context.library, context.topic, context.key)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(rec.Value)
		}
	} else {
		fmt.Println("Must be in path with library/topic/key")
	}

}