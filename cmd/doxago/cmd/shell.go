// Copyright © 2020 The Orthodox Christian Mission Center (ocmc.org)

package cmd

import (
	SQL "database/sql"
	"encoding/json"
	"fmt"
	prompt "github.com/c-bata/go-prompt"
	"github.com/liturgiko/doxa/pkg/concord"
	"github.com/liturgiko/doxa/pkg/db/ltx2sql"
	"github.com/liturgiko/doxa/pkg/models"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"log"
	"os"
	"strconv"
	"strings"
)

// the Shell command provides a user-friendly means search liturgical texts and create of update values
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

		settings.Padding.P1 = 4
		settings.Padding.P2 = 50
		settings.Padding.P3 = 0
		settings.ShowEmpty = false
		settings.Sort = concord.SortRight
		settings.Width = 30

		// Get the width and height of the terminal.
		// We also do this whenever find is called in
		// case it got resized.
		tWidth, tHeight, err := terminal.GetSize(0)
		if err != nil {
			fmt.Println(err)
		} else {
			context.THeight = tHeight
			context.TWidth = tWidth
		}

		setSuggestions()
		p := prompt.New(
			executor,
			completer,
			prompt.OptionPrefix(context.Path+"> "),
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
var cMap map[string]concord.ConcordanceLine
var suggestions []prompt.Suggest
var pathDelimiter = "/" // strconv.QuoteRune(os.PathSeparator)
type Padding struct {
	P1 int `json:"number"`
	P2 int `json:"id"`
	P3 int `json:"value"`
}
type Settings struct {
	Exact     bool          `json:".exact"`
	IDlike    string        `json:".idlike"`
	Padding   Padding       `json:".padding"`
	Vfirst    bool          `json:".vfirst"`
	Width     int           `json:".width"`
	ShowEmpty bool          `json:".showempty"`
	Sort      concord.Order `json:".sort"`
}

var settings Settings

type Context struct {
	Library string // library of current path, e.g. gr_gr_cog
	Topic   string // topic of current path, e.g. actors
	Key     string // key of current path, e.g. Priest
	Path    string // current path, e.g. gr_gr_cog/actors/Priest
	THeight int    // height of the user's terminal window
	TWidth  int    // Width of the user's terminal window
}

var context Context
var previousContext Context

func (c *Context) DBPath() string {
	id := models.Id{}
	id.Set(c.Library, c.Topic, c.Key)
	return id.ToNeoId()
}
func (c *Context) SetPath() {
	sb := strings.Builder{}
	if len(c.Library) > 0 {
		sb.WriteString(c.Library)
		if len(c.Topic) > 0 {
			sb.WriteString(pathDelimiter)
			sb.WriteString(c.Topic)
			if len(c.Key) > 0 {
				sb.WriteString(pathDelimiter)
				sb.WriteString(c.Key)
			}
		}
		context.Path = sb.String()
	}
}
func (c *Context) Depth() int {
	if context.Path == "" {
		return 0
	} else {
		return len(strings.Split(context.Path, pathDelimiter))
	}
}

// Based on the depth of the path, provides a LIKE clause for sql select where like
func (c *Context) Like() string {
	like := ""
	switch c.Depth() {
	case 1:
		like = c.Library + "~%"
	case 2:
		like = c.Library + "~" + c.Topic + "~%"
	case 3:
		like = c.Library + "~" + c.Topic + "~" + c.Key
	}
	return like
}
func (c *Context) GetPath() string {
	var b strings.Builder
	if len(context.Key) > 0 {
		b.WriteString(context.Library)
		b.WriteString(pathDelimiter)
		b.WriteString(context.Topic)
		b.WriteString(pathDelimiter)
		b.WriteString(context.Key)
		return b.String()
	}
	if len(context.Topic) > 0 {
		b.WriteString(context.Library)
		b.WriteString(pathDelimiter)
		b.WriteString(context.Topic)
		return b.String()
	}
	return context.Library
}

// The map help by this type provides occurrence numbers
// as aliases for IDs to make it easy for the user to cd to a path.
type IDMap struct {
	Map map[int]models.Id
}

func (im *IDMap) Reset() {
	im.Map = make(map[int]models.Id)
}
func (im *IDMap) Add(index int, path string) {
	var id models.Id
	id.Parse(strings.TrimSpace(path))
	im.Map[index] = id
}

// This map provides occurrence numbers as aliases for IDs
// to make it easy for the user to cd to a path.
var idMap IDMap

func livePrefix() (string, bool) {
	if context.Path == "" {
		return "", false
	}
	return context.Path + "> ", true
}

// executes the commands issued by the user
func executor(in string) {
	in = strings.TrimSpace(in)
	var method string
	blocks := strings.Split(in, " ")
	switch blocks[0] {
	case "cd":
		copyContext()
		if len(blocks) < 2 {
			context.Path = ""
		} else if blocks[1] == "~" {
			context.Path = resetPath()
		} else if len(blocks) > 1 && isNumber(blocks[1]) {
			index, err := strconv.Atoi(blocks[1])
			if err != nil {
				fmt.Println(err)
			} else {

				if len(idMap.Map[index].Domain.Country) > 0 {
					setContextId(idMap.Map[index])
				} else {
					fmt.Printf("%s not found\n", blocks[1])
				}
			}
		} else {
			// For user convenience, if the user pasted a DB ID with tildes
			// convert to /.  Because an initial tilde is an alias for $HOME
			// we skip the first char when doing the replacement.
			var path string
			if strings.HasPrefix(blocks[1], "~") {
				path = "~" + strings.ReplaceAll(blocks[1][1:], "~", "/")
			} else {
				path = strings.ReplaceAll(blocks[1], "~", "/")
			}
			up, dirs := cdParts(path)
			if len(up) > 0 {
				if up == "~" {
					context.Path = resetPath()
				} else {
					context.Path = popPath(up)
				}
			}
			if len(dirs) > 0 {
				s := strings.Split(dirs, pathDelimiter)
				depth := context.Depth()
				if depth+len(s) > 3 {
					fmt.Println("You can't go deeper than 3 levels (library/topic/key).")
				} else {
					for _, segment := range s {
						context.Path = pushPath(segment)
					}
				}
			}
			if context.Depth() == 3 {
				if exists(context.Library, context.Topic, context.Key) {
					showValue(context.Library, context.Topic, context.Key)
				} else {
					revertContext()
					fmt.Printf("%s not found\n", blocks[1])
				}
			} else {
				reportCount()
			}
		}
		setSuggestions()
		return
	case "cmp":
		{
			method = "cmp"
			idMap.Reset()
			if context.Depth() == 3 {
				recs, err := mapper.ReadByTK(context.Topic, context.Key, settings.ShowEmpty)
				if err != nil {
					fmt.Println(err)
					return
				}
				for i, rec := range recs {
					idMap.Add(i+1, rec.ID)
					fmt.Printf("%s\n", strings.Repeat("-", context.TWidth))
					fmt.Printf("%4d %s", i+1, rec.ID)
					fmt.Printf("%s\n", strings.Repeat("-", context.TWidth))
					fmt.Println(rec.Value)
				}
			} else {
				fmt.Println("You must be three levels deep to use this command.")
			}
		}
	case "find":
		{
			method = "find"
			cMap = make(map[string]concord.ConcordanceLine)
			idMap.Reset()
			tWidth, tHeight, err := terminal.GetSize(0)
			if err != nil {
				fmt.Println(err)
			} else {
				context.THeight = tHeight
				context.TWidth = tWidth
			}

			if len(blocks) == 1 {
				fmt.Println("You must tell me what value to find")
			} else {
				value := escape(strings.Join(blocks[1:len(blocks)], " "))
				if strings.Contains(value, "%") {
					// at this time, it is unlikely that there is a % in liturgical text.
					// But, just in case, we will escape the percent
					value = strings.ReplaceAll(value, "%", "\\%")
				}
				var recs []*models.Ltx
				var err error
				var id string
				if len(settings.IDlike) > 0 {
					id = settings.IDlike
				} else {
					var sb strings.Builder
					switch context.Depth() {
					case 1:
						sb.WriteString(context.Library)
						sb.WriteString("~%")
					case 2:
						sb.WriteString(context.Library)
						sb.WriteString("~")
						sb.WriteString(context.Topic)
						sb.WriteString("~%")
					case 3:
						sb.WriteString(context.Library)
						sb.WriteString("~")
						sb.WriteString(context.Topic)
						sb.WriteString("~")
						sb.WriteString(context.Key)
					}
					id = sb.String()
				}
				if settings.Exact {
					recs, err = mapper.ReadByValue(id, value)
				} else {
					recs, err = mapper.ReadByNNP(id, value)
				}
				if err != nil {
					fmt.Println(err)
				} else {
					for _, rec := range recs {
						var line = ""
						if settings.Exact {
							line = rec.Value
						} else {
							line = rec.NNP
						}

						cline(fmt.Sprintf("%*s", settings.Padding.P2, rec.ID), line, value, settings.Width)
					}
					for i, res := range concord.SortedKeys(cMap, settings.Sort) {
						idMap.Add(i+1, cMap[res].ID)
						if len(strings.TrimSpace(cMap[res].ID)) > 50 {
							parts := strings.Split(cMap[res].ID, "~")
							if len(parts) == 3 {
								fmt.Printf("%*d |%s~%s~\n", settings.Padding.P1, i+1, parts[0], parts[1])
								fmt.Printf("     | %*s | %-s%s%s\n", settings.Padding.P2, parts[2], cMap[res].Left, cMap[res].Key, cMap[res].Right)
							} else {
								// bad
								fmt.Printf("%*d | %*s | %-s%s%s\n", settings.Padding.P1, i+1, settings.Padding.P2, cMap[res].ID, cMap[res].Left, cMap[res].Key, cMap[res].Right)
							}
						} else {
							fmt.Printf("%*d | %*s | %-s%s%s\n", settings.Padding.P1, i+1, settings.Padding.P2, cMap[res].ID, cMap[res].Left, cMap[res].Key, cMap[res].Right)
						}
					}
					var exact = "off"
					if settings.Exact {
						exact = "off"
					}
					fmt.Printf("%d records found. Find .exact %s. To change: .exact on or .exact off", len(recs), exact)
					if len(settings.IDlike) > 0 {
						fmt.Printf(".idlike = %s, so current path was not used. To turn off: idlike %%", settings.IDlike)
					}
					fmt.Println("")

				}
			}
		}
	case "lml":
		if context.Depth() == 3 {
			msg := fmt.Sprintf(`"%s~%s"`,context.Topic, context.Key)
			fmt.Println(msg)
		} else {
			fmt.Println("You must be three levels deep to use this command")
		}
		return
	case "ls":
		method = "ls"
		if len(blocks) > 1 && blocks[1] == "empty" {
			idMap.Reset()
			like := context.Like()
			ids, err := mapper.Empty(like)
			if err != nil {
				Logger.Print(err)
			}
			for i, id := range ids {
				idMap.Add(i+1, id)
				fmt.Printf("%4d %s\n", i+1, id)
			}
			fmt.Printf("%d empty records found\n",len(ids))
		} else if len(blocks) > 1 && blocks[1] == "redirects" {
			idMap.Reset()
			like := context.Like()
			if context.Depth() < 3 {
				redirects, err := mapper.Redirects(like)
				if err != nil {
					Logger.Print(err)
				}
				for i, redirect := range redirects {
					idMap.Add(i+1, redirect.ID)
					fmt.Printf("%4d %s ==> %s\n", i+1, redirect.ID, redirect.Redirect)
				}
			} else {
				redirects, err := mapper.ReferredTo(like)
				if err != nil {
					Logger.Print(err)
				}
				for i, redirect := range redirects {
					idMap.Add(i+1, redirect.Redirect)
					fmt.Printf("%4d %s <== %s\n", i+1, redirect.Redirect, redirect.ID)
				}
			}
		} else if len(blocks) > 1 && isNumber(blocks[1]) {
			index, err := strconv.Atoi(blocks[1])
			if err != nil {
				fmt.Println(err)
			} else {
				if idMap.Map[index].HasValues() {
					id := idMap.Map[index]
					showValue(id.Domain.ToNeo(), id.Topic, id.Key)
				} else {
					fmt.Printf("%s not found\n", blocks[1])
				}
			}
		} else {
			if context.Depth() < 3 {
				idMap.Reset()
			}
			switch context.Depth() {
			case 0:
				{ // list all libraries in database
					libraries, err := mapper.Libraries()
					if err != nil {
						Logger.Print(err)
					}
					for i, library := range libraries {
						var id models.Id
						var d models.Domain
						d.Parse(library)
						id.Domain = d
						idMap.Add(i+1, id.ToNeoId())
						fmt.Printf("%4d %s\n", i+1, library)
					}
				}
			case 1:
				{ // list topics for current library
					var like = context.Library
					if len(blocks) > 1 {
						like = like + "~" + blocks[1]
					}
					topics, err := mapper.Topics(like)
					if err != nil {
						Logger.Print(err)
					}
					for i, topic := range topics {
						var id models.Id
						var d models.Domain
						d.Parse(context.Library)
						id.Domain = d
						id.Topic = topic
						idMap.Add(i+1, id.ToNeoId())
						fmt.Printf("%4d %s\n", i+1, topic)
					}
				}
			case 2:
				{ // list keys for current library/topic
					var like = context.Library + "~" + context.Topic
					if len(blocks) > 1 {
						like = like + "~" + blocks[1]
					}
					keys, err := mapper.Keys(like)
					if err != nil {
						Logger.Print(err)
					}
					for i, key := range keys {
						var id models.Id
						var d models.Domain
						d.Parse(context.Library)
						id.Domain = d
						id.Topic = context.Topic
						id.Key = key
						idMap.Add(i+1, id.ToNeoId())
						fmt.Printf("%4d %s\n", i+1, key)
					}
				}
			case 3:
				{ // Display the value of the key
					showValue(context.Library, context.Topic, context.Key)
				}
			}
		}
	case "set":
		// TODO: when ltx struct is modified to have properties for create/modify date and by whom, need to set these here.
		method = "set"
		if context.Depth() < 3 {
			fmt.Println("You must be three levels deep to use the set command")
			return
		}
		rec, err := mapper.ReadByLTK(context.Library, context.Topic, context.Key)
		if err != nil {
			fmt.Println(err)
		} else {
			var prompt = fmt.Sprintf("set what, to what?\nExample 1: set value PRIEST\nExample 2: set redirect gr_gr_cog/template.titles/d.onSaturdayEvening\nTo set blank: set value ''\nWhen setting to blank, use '' not \"\"")
			value := ""
			if len(blocks) > 1 {
				if len(blocks) > 2 {
					value = escape(strings.Join(blocks[2:len(blocks)], " "))
					switch blocks[1] {
					case "redirect":
						if value == "''" {
							value = ""
							rec.SetRedirect(value)
						} else {
							value = strings.ReplaceAll(value, "/", "~") // in case user used / instead of ~
							var id models.Id
							id.Parse(value)
							if exists(id.Domain.ToNeo(), id.Topic, id.Key) {
								rec.SetRedirect(value)
							} else {
								fmt.Printf("Does not exist in database: %s\n", value)
							}
						}
					case "value":
						if value == "''" {
							value = ""
						}
						rec.SetValue(value)
					default:
						fmt.Println(prompt)
					}
					mapper.Merge(rec)
					if len(rec.Value) > 0 {
						fmt.Println(rec.Value)
					} else {
						if len(rec.Redirect) > 0 {
							idMap.Reset()
							idMap.Add(1, value)
							fmt.Printf("Redirects to: 1 %s\n", rec.Redirect)
						}
					}
				} else {
					fmt.Println(prompt)
				}
			} else {
				fmt.Println(prompt)
			}
		}
	case ".context":
		{
			method = ".context"
			var jsonData []byte
			jsonData, err := json.Marshal(context)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(string(jsonData))
		}
	case ".exact":
		{
			method = ".exact"
			if len(blocks) > 1 {
				settings.Exact = strings.Contains(strings.ToLower(blocks[1]), "on")
				if settings.Exact {
					err := mapper.CaseSensitiveLike(true)
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println("on")
				} else {
					err := mapper.CaseSensitiveLike(false)
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println("off")
				}
			} else {
				if settings.Exact {
					fmt.Println("Exact match for find is on. To switch off: .exact off")
				} else {
					fmt.Println("Exact match for find is off. To switch on: .exact on")
				}
			}
		}
	case ".exit":
		fmt.Println("Bye!")
		os.Exit(0)
	case ".idlike":
		{
			method = ".idlike"
			if len(blocks) > 1 {
				if blocks[1] == "%" {
					settings.IDlike = ""
					fmt.Println(".idlike turned off")
				} else {
					settings.IDlike = blocks[1]
					if !strings.HasPrefix(settings.IDlike, "%") {
						settings.IDlike = "%" + settings.IDlike
					}
					if !strings.HasSuffix(settings.IDlike, "%") {
						settings.IDlike = settings.IDlike + "%"
						fmt.Println(settings.IDlike)
						fmt.Println("To turn off: .idlike %")
					}
				}
			} else {
				fmt.Printf("You must provide an id pattern, e.g. en%% or %%dedes%%\n")
			}
		}
	case ".padding":
		{
			method = ".padding"
			switch len(blocks) {
			case 1:
				{
					fmt.Println("number |id                        |value")
					fmt.Println("2433   |en_us_net~ps~psa44.v8.text| anointed you with the oil of joy elevating you above your co")
					fmt.Println("The `Padding` command controls the Padding of the three parts of a concordance line shown as the result of the find command. The first Padding is for the result number. The second and third paddings are for the id and value (or vice versa if Vfirst = on). Padding values can be negative, in which case they are left aligned.  Positive values are right aligned.")
					fmt.Printf("The current values are: %v\n", settings.Padding)
				}
			case 2:
				{
					setPadding(1, blocks[1])
				}
			case 3:
				{
					setPadding(1, blocks[1])
					setPadding(2, blocks[2])
				}
			case 4:
				{
					setPadding(1, blocks[1])
					setPadding(2, blocks[2])
					setPadding(3, blocks[3])
				}
			}
		}
	case ".showempty":
		{
			method = ".showempty"
			if len(blocks) > 1 {
				settings.ShowEmpty = strings.Contains(strings.ToLower(blocks[1]), "on")
				if settings.ShowEmpty {
					err := mapper.CaseSensitiveLike(true)
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println("on")
				} else {
					err := mapper.CaseSensitiveLike(false)
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println("off")
				}
			} else {
				if settings.ShowEmpty {
					fmt.Println(".showempty is on. The `find` command will include records with an empty value. To switch off: .showempty off")
				} else {
					fmt.Println(".showempty is off. The `find` command will exclude records with an empty value. To switch on: .showempty on")
				}
			}
		}
	case ".vfirst":
		{
			method = ".vfirst"
			if len(blocks) > 1 {
				settings.Vfirst = strings.Contains(strings.ToLower(blocks[1]), "on")
				if settings.Vfirst {
					fmt.Println("on")
				} else {
					fmt.Println("off")
				}
			} else {
				if settings.Vfirst {
					fmt.Println(".vfirst controls the output of the find command. `on` shows value, then id. `off` shows id, then value. match for find is on. To switch off: .vfirst off")
				} else {
					fmt.Println("Show value first for find is off. To switch on: .vfirst on")
				}
			}
		}
	case ".width":
		{
			method = ".width"
			switch len(blocks) {
			case 2:
				{
					i, err := strconv.Atoi(blocks[1])
					if err != nil {
						fmt.Println(err)
					} else {
						settings.Width = i
					}
				}
			default:
				{
					fmt.Println("`Width` sets the Width of the value shown using the `find` command.")
					fmt.Printf("Width = %d\n", settings.Width)
				}
			}
		}
	case ".settings":
		{
			method = ".settings"
			var jsonData []byte
			jsonData, err := json.Marshal(settings)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(string(jsonData))
		}
	case ".sort":
		{
			method = ".sort"
			if len(blocks) > 1 {
				order := strings.ToLower(blocks[1])
				switch order {
				case "id":
					settings.Sort = concord.SortId
					fmt.Println("Sort by id is on")
				case "left":
					settings.Sort = concord.SortLeft
					fmt.Println("Sort by left is on")
				case "right":
					settings.Sort = concord.SortRight
					fmt.Println("Sort by right is on")
				default:
					fmt.Println("Invalid sort option. Use: .sort id or .sort left or .sort right")
				}
			} else {
				fmt.Printf("%s. To change: .sort id or .sort left or .sort right\n", settings.Sort.String())
			}
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
	s := len(strings.Split(context.Path, pathDelimiter))
	r := s - l
	if r < 1 {
		context.Library = ""
		context.Topic = ""
		context.Key = ""
	} else if r == 1 {
		context.Topic = ""
		context.Key = ""
	} else if r == 2 {
		context.Key = ""
	} else {
		fmt.Println("Oops!")
	}
	return context.GetPath()
}
func pushPath(segment string) string {
	if len(context.Library) < 1 {
		context.Library = segment
	} else if len(context.Topic) < 1 {
		context.Topic = segment
	} else {
		context.Key = segment
	}
	return context.GetPath()
}
func setSuggestions() {
	suggestions = []prompt.Suggest{
		// Command
		{"cd", "CHANGE path"},
		{"cp", "COPY contents"},
		{"Exact", "EXACT on, EXACT off. If on, find will match case and accents."},
		{"exit", "Exit Doxa Shell"},
		{"find", "FIND records with specified value"},
		{"get", "GET value for specified id, e.g. get en_us_dedes~actors~Priest"},
		{"ls", "LIST contents"},
		{"mv", "MOVE (rename) matching id to new id"},
		{"rm", "REMOVE for matching id"},
		{"set", "SET value for specified id, e.g. set en_us_dedes~actors~Priest=\"Priest\""},
	}
	switch context.Depth() {
	case 0:
		appendLibraries()
	case 1:
		appendTopics()
	default:
		appendKeys()
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
	context.Library = ""
	context.Topic = ""
	context.Key = ""
	return context.GetPath()
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
func exists(library, topic, key string) bool {
	c, _ := mapper.CountKeys(library, topic, key)
	return c == 1
}
func reportCount() {
	if len(context.Path) > 0 {
		p := message.NewPrinter(language.English)
		c, _ := mapper.CountKeys(context.Library, context.Topic, context.Key)
		if c == 0 {
			fmt.Printf("Path %s does not exist\n", context.Path)
			revertContext()
		} else {
			if context.Depth() == 1 {
				t, _ := mapper.CountTopics(context.Library)
				p.Printf("%d topics %d keys\n", t, c)
			} else {
				p.Printf("%d keys\n", c)
			}
		}
	}
}
func copyContext() {
	previousContext.Library = context.Library
	previousContext.Topic = context.Topic
	previousContext.Key = context.Key
	previousContext.Path = context.Path
}
func revertContext() {
	context.Library = previousContext.Library
	context.Topic = previousContext.Topic
	context.Key = previousContext.Key
	context.Path = previousContext.Path
}
func showValue(library, topic, key string) {
	rec, err := mapper.ReadByLTK(library, topic, key)
	if err != nil {
		fmt.Println(err)
	} else {
		if rec != nil {
			if len(rec.Redirect) > 0 {
				fmt.Printf("Redirects to: 1 %s\n", rec.Redirect)
				var id models.Id
				err = id.Parse(rec.Redirect)
				if err != nil {
					fmt.Println(err)
					return
				}
				idMap.Reset()
				idMap.Add(1, rec.Redirect)
				showValue(id.Domain.ToNeo(), id.Topic, id.Key)
			} else {
				fmt.Println(rec.Value)
			}
		}
	}
}
func escape(value string) string {
	value = strconv.Quote(value)
	// a side effect of strconv.Quote is it wraps the entire string in quotes,
	// which we do not want.  So, remove them.
	value = value[1 : len(value)-1]
	return value
}

// Centers keyword in middle, with window size = Width for left and right
// And, add id, left, key, right to concordance map named cMap so we can Sort the lines
// by parts.
func cline(id, line, key string, width int) string {
	r := []rune(line)
	rKey := []rune(key)
	keyIndex := indexInRune(r, key)
	lineLen := len(r)
	// get left context
	leftIndex := 0
	if keyIndex-width > 0 {
		leftIndex = keyIndex
		for i := keyIndex - 1; i > 0; i-- {
			if leftIndex <= (keyIndex - width) {
				break
			} else {
				leftIndex--
			}
		}
	}
	// get right context
	rightStart := keyIndex + len(rKey) + 1
	rightIndex := rightStart
	for i := rightStart; i < lineLen; i++ {
		if rightIndex > keyIndex+width {
			break
		} else {
			rightIndex++
		}
	}
	var left, right string
	if 0 <= leftIndex && leftIndex <= keyIndex && keyIndex <= len(r) {
		left = string(r[leftIndex:keyIndex])
	}
	left = fmt.Sprintf("%*s", width-len(left), left)
	keyRight := keyIndex + len(rKey)
	if keyRight < rightIndex && rightIndex <= len(r) {
		right = string(r[keyRight:rightIndex])
	} else {
		if keyRight < len(r) {
			right = string(r[keyRight:])
		}
	}
	right = fmt.Sprintf("%-*s", width-len(right), right)
	var cl concord.ConcordanceLine
	cl.ID = id
	cl.Left = fmt.Sprintf("%*s", width+3, left)
	cl.Key = key
	cl.Right = right
	cMap[cl.ID] = cl
	return fmt.Sprintf("%*s%s%s", width+3, left, key, right)
}

// Find the index of a substring within []rune
// This solves the following problem:
// When we have a string of Greek, they are runes. If we take a slice,
// we will get ? output for some characters.  The solution is to convert the
// string to a []rune.  But, then, our search index won't work since the length
// of the original string is > than the length of the []rune.
func indexInRune(text []rune, what string) int {
	whatRunes := []rune(what)
	for i := range text {
		found := true
		for j := range whatRunes {
			if text[i+j] != whatRunes[j] {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}
	return -1
}

// converts parameter of paddings to int and stores in settings.Padding for indicated index of 1..3
func setPadding(index int, value string) {
	i, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch index {
	case 1:
		settings.Padding.P1 = i
	case 2:
		settings.Padding.P2 = i
	case 3:
		settings.Padding.P3 = i
	default:
		fmt.Printf("Invalid index %d passed to func setPadding\n", index)
	}
}
func isNumber(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	} else {
		return false
	}
}
func setContextId(id models.Id) {
	if id.HasValues() {
		context.Library = id.Domain.ToNeo()
		context.Topic = id.Topic
		context.Key = id.Key
		context.SetPath()
	} else {
		fmt.Printf("setContextId parse error. Expected library, topic, key\n")
	}
}
