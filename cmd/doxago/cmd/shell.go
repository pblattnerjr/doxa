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
var cMap map[string]concord.ConcordanceLine
var suggestions []prompt.Suggest
var pathDelimiter = "/" // strconv.QuoteRune(os.PathSeparator)
type Padding struct {
	P1 int `json:"number"`
	P2 int `json:"id"`
	P3 int `json:"value"`
}
type Settings struct {
	Exact   bool          `json:".exact"`
	IDlike  string        `json:".idlike"`
	Padding Padding       `json:".padding"`
	Vfirst  bool          `json:".vfirst"`
	Width   int           `json:".width"`
	ShowEmpty bool `json:".showempty"`
	Sort    concord.Order `json:".sort"`
}

var settings Settings

type Context struct {
	library string // library of current path, e.g. gr_gr_cog
	topic   string // topic of current path, e.g. actors
	key     string // key of current path, e.g. Priest
	path    string // current path, e.g. gr_gr_cog/actors/Priest
	tHeight int    // height of the user's terminal window
	tWidth  int    // Width of the user's terminal window
}

var context Context
var previousContext Context

func (c *Context) DBPath() string {
	id := models.Id{}
	id.Set(c.library, c.topic, c.key)
	return id.ToNeoId()
}
func (c *Context) SetPath() {
	sb := strings.Builder{}
	if len(c.library) > 0 {
		sb.WriteString(c.library)
		if len(c.topic) > 0 {
			sb.WriteString(pathDelimiter)
			sb.WriteString(c.topic)
			if len(c.key) > 0 {
				sb.WriteString(pathDelimiter)
				sb.WriteString(c.key)
			}
		}
		context.path = sb.String()
	}
}
func (c *Context) Depth() int {
	if context.path == "" {
		return 0
	} else {
		return len(strings.Split(context.path, pathDelimiter))
	}
}

// Based on the depth of the path, provides a LIKE clause for sql select where like
func (c *Context) Like() string {
	like := ""
	switch c.Depth() {
	case 1:
		like = c.library + "~%"
	case 2:
		like = c.library + "~" + c.topic + "~%"
	case 3:
		like = c.library + "~" + c.topic + "~" + c.key
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

var idMap map[int]models.Id // holds IDs of matching records resulting from ls or find

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
	case "cd":
		copyContext()
		if len(blocks) < 2 {
			context.path = ""
		} else if blocks[1] == "~" {
			context.path = resetPath()
		} else if len(blocks) > 1 && isNumber(blocks[1]) {
			index, err := strconv.Atoi(blocks[1])
			if err != nil {
				fmt.Println(err)
			} else {
				if len(idMap[index].Domain.Country) > 0 {
					setContextId(idMap[index])
				} else {
					fmt.Printf("%s not found\n", blocks[1])
				}
			}
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
				if depth+len(s) > 3 {
					fmt.Println("You can't go deeper than 3 levels (library/topic/key).")
				} else {
					for _, segment := range s {
						context.path = pushPath(segment)
					}
				}
			}
			if context.Depth() == 3 {
				if exists(context.library, context.topic, context.key) {
					showValue(context.library, context.topic, context.key)
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
			if context.Depth() == 3 {
				recs,err := mapper.ReadByTK(context.topic, context.key,settings.ShowEmpty)
				if err != nil {
					fmt.Println(err)
					return
				}
				for _,rec := range recs {
					fmt.Printf("%s\n",strings.Repeat("-",context.tWidth))
					fmt.Println(rec.ID)
					fmt.Printf("%s\n",strings.Repeat("-",context.tWidth))
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
			idMap = make(map[int]models.Id)
			tWidth, tHeight, err := terminal.GetSize(0)
			if err != nil {
				fmt.Println(err)
			} else {
				context.tHeight = tHeight
				context.tWidth = tWidth
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
						sb.WriteString(context.library)
						sb.WriteString("~%")
					case 2:
						sb.WriteString(context.library)
						sb.WriteString("~")
						sb.WriteString(context.topic)
						sb.WriteString("~%")
					case 3:
						sb.WriteString(context.library)
						sb.WriteString("~")
						sb.WriteString(context.topic)
						sb.WriteString("~")
						sb.WriteString(context.key)
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
						var id models.Id
						id.Parse(strings.TrimSpace(cMap[res].ID))
						idMap[i+1] = id
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
	case "ls":
		method = strings.ToUpper(blocks[0])
		if len(blocks) > 1 && isNumber(blocks[1]) {
			index, err := strconv.Atoi(blocks[1])
			if err != nil {
				fmt.Println(err)
			} else {
				if idMap[index].HasValues() {
					id := idMap[index]
					showValue(id.Domain.ToNeo(), id.Topic, id.Key)
				} else {
					fmt.Printf("%s not found\n", blocks[1])
				}
			}
		} else {
			idMap = make(map[int]models.Id)
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
						idMap[i+1] = id
						fmt.Printf("%4d %s\n", i+1, library)
					}
				}
			case 1:
				{ // list topics for current library
					var like = context.library
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
						d.Parse(context.library)
						id.Domain = d
						id.Topic = topic
						idMap[i+1] = id
						fmt.Printf("%4d %s\n", i+1, topic)
					}
				}
			case 2:
				{ // list keys for current library/topic
					var like = context.library + "~" + context.topic
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
						d.Parse(context.library)
						id.Domain = d
						id.Topic = context.topic
						id.Key = key
						idMap[i+1] = id
						fmt.Printf("%4d %s\n", i+1, key)
					}
				}
			case 3:
				{ // Display the value of the key
					showValue(context.library, context.topic, context.key)
				}
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
				value = escape(strings.Join(blocks[1:len(blocks)], " "))
			}
			rec.SetValue(value)
			mapper.Merge(rec)
			fmt.Println(rec.Value)
		}
	case ".context":
		{
			method = ".context"
			fmt.Printf("Library: %s, Topic: %s, Key: %s, Path: %s, Depth: %d\n", context.library, context.topic, context.key, context.path, context.Depth())
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
func exists(library, topic, key string) bool {
	c, _ := mapper.CountKeys(context.library, context.topic, context.key)
	return c == 1
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
				p.Printf("%d topics %d keys\n", t, c)
			} else {
				p.Printf("%d keys\n", c)
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
func showValue(library, topic, key string) {
	rec, err := mapper.ReadByLTK(library, topic, key)
	if err != nil {
		fmt.Println(err)
	} else {
		if rec != nil {
			fmt.Println(rec.Value)
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
		context.library = id.Domain.ToNeo()
		context.topic = id.Topic
		context.key = id.Key
		context.SetPath()
	} else {
		fmt.Printf("setContextId parse error. Expected library, topic, key\n")
	}
}
