// Copyright © 2020 The Orthodox Christian Mission Center (ocmc.org)

package cmd

import (
	SQL "database/sql"
	"encoding/json"
	"fmt"
	"github.com/c-bata/go-prompt"
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

// the Shell command provides a user-friendly means to search liturgical texts and create or update values
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "provide a shell for accessing the liturgical database",
	Long:  `provide a shell for accessing the liturgical database`,
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
var conc concord.Concordance
var commands []prompt.Suggest
var suggestions []prompt.Suggest
var pathDelimiter = mapper.IDDelimiter()

type Padding struct {
	P1 int `json:"number"`
	P2 int `json:"id"`
	P3 int `json:"value"`
}
type Settings struct {
	Exact     bool          `json:".exact"`
	Hints     bool          `json:".hints"`
	IDlike    string        `json:".idlike"`
	Padding   Padding       `json:".padding"`
	Width     int           `json:".width"`
	ShowAll   bool          `json:".showall"`
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
	return id.ToSQLId()
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

// Based on the depth of the path, provides a sql SELECT WHERE LIKE clause
func (c *Context) Like() string {
	like := ""
	switch c.Depth() {
	case 1:
		like = c.Library + pathDelimiter + "%"
	case 2:
		like = c.Library + pathDelimiter + c.Topic + pathDelimiter + "%"
	case 3:
		like = c.Library + pathDelimiter + c.Topic + pathDelimiter + c.Key
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

// The map held by this type provides occurrence numbers
// as aliases for IDs to make it easy for the user to cd to a path.
type IDMap struct {
	Map map[int]models.Id
}

// Resets the map so it is empty
func (im *IDMap) Reset() {
	im.Map = make(map[int]models.Id)
}

// Adds a path to the map
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
		method = blocks[0]
		changeDirectory(blocks)
	case "cmp":
		method = blocks[0]
		compareValues()
	case "exit":
		fmt.Println("Bye!")
		os.Exit(0)
	case "find":
		method = blocks[0]
		find(blocks)
	case "help":
		method = blocks[0]
		showHelp()
	case "lml":
		method = blocks[0]
		showTKAsLML()
	case "ls":
		method = blocks[0]
		list(blocks)
	case "set":
		// TODO: when ltx struct is modified to have properties for create/modify by whom, need to set these here.
		method = blocks[0]
		setRecord(blocks)
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
	case ".hints":
		{
			method = ".hints"
			if len(blocks) > 1 {
				settings.Hints = strings.Contains(strings.ToLower(blocks[1]), "on")
				if settings.Hints {
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
				if settings.Hints {
					fmt.Println("Hints are on. To switch off: .hints off")
				} else {
					fmt.Println("Exact match for find is off. To switch on: .hints on")
				}
			}
		}
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
					fmt.Println("2433   |en_us_net/ps/psa44.v8.text| anointed you with the oil of joy elevating you above your co")
					fmt.Println("The `Padding` command controls the Padding of the three parts of a concordance line shown as the result of the find command. The first Padding is for the result number. The second and third paddings are for the id and value). Padding values can be negative, in which case they are left aligned.  Positive values are right aligned.")
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
	case ".showall":
		{
			method = ".showall"
			if len(blocks) > 1 {
				settings.ShowAll = strings.Contains(strings.ToLower(blocks[1]), "on")
				if settings.ShowAll {
					fmt.Println("on")
				} else {
					fmt.Println("off")
				}
			} else {
				if settings.ShowEmpty {
					fmt.Println(".showall is on. When 3 levels deep, ls will show the entire record. To switch off: .showall off")
				} else {
					fmt.Println(".showall is off. When 3 levels deep, ls will only show the value. To see all record properties: .showall on")
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
					fmt.Println("`Width` sets the Width of left and right parts of the concordance lines shown using the `find` command.")
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

	fmt.Println("Sorry, I don't understand.")
}

func completer(in prompt.Document) []prompt.Suggest {
	w := in.GetWordBeforeCursor()
	if w == "" {
		return []prompt.Suggest{}
	}
	return prompt.FilterHasPrefix(suggestions, w, true)
}

// moves up the path hierarchy
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

// Moves deeper down the path hierarchy
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

// Sets the values for the popup window that shows the users matching strings
func setSuggestions() {
	suggestions = []prompt.Suggest{}
	appendCommands()
	switch context.Depth() {
	case 0:
		appendLibraries()
	case 1:
		appendTopics()
	default:
		appendKeys()
	}
}

// add commands to suggestions
func appendCommands() {
	commands = []prompt.Suggest{
		// Command
		{"cd", "CHANGE path, e.g. cd gr_gr_cog/actors or cd gr_gr_cog/actors/Priest"},
		{"cd ..", "CHANGE path up one level"},
		{"cd ../..", "CHANGE path up two levels"},
		{"cd ~", "CHANGE path to root"},
		{"cd {number}", "After a find or ls, numbers can be used to change path, e.g. cd 244"},
		{"cmp", "compare values for this topic/key. Must be 3 levels deep."},
		{"cp", "*COPY contents"},
		{"exit", "Exit Doxago Shell"},
		{"find", "FIND records with specified value"},
		{"lml", "Display topic/key in the format required for the Liturgical Markup Language."},
		{"ls", "LIST contents. If at root, lists libraries. If 3 levels deep shows record value"},
		{"mk", "*MAKE. At root, mk en_us_xyz makes library. 1 level deep, mk actors makes new topic, 3 deep mk Priest makes new key"},
		{"mv", "*MOVE (rename) matching id to new id"},
		{"rm", "*REMOVE for matching id"},
		{"set comment", "SET comment for current record. Must be 3 levels deep."},
		{"set redirect", "SET redirect for current record. Must be 3 levels deep."},
		{"set value", "SET value for current record. Must be 3 levels deep."},
		{".context", "Shows the current context, i.e. the current path"},
		{".exact off", "If EXACT off, find is insensitive to case and accents and punctuation."},
		{".exact on", "If EXACT on, find will match case and accents."},
		{".hints off", "If HINTS off, no explanations will be given about what you can do next."},
		{".hints on", "If HINTS on, explanations of what you can do next will be displayed."},
		{".idlike", "Sets pattern of id for subsequent finds, e.g. .idlike en will only match ids starting with en"},
		{".padding", "Sets the padding for concordance line parts."},
		{".settings", "Display the values of the settings"},
		{".showall off", "Show only the value of the record"},
		{".showall on", "Show the entire record"},
		{".showempty", "Show records where both the value and redirect are empty"},
		{".showempty on", "Show the entire record"},
		{".sort id", "Results of Find will be sorted by record ID"},
		{".sort left", "Results of Find will be sorted by left part of concordance line"},
		{".sort right", "Results of Find will be sorted by right part of concordance line"},
		{".width", "Sets the width of left and right sides of concordance line"},
	}
	for _, command := range commands {
		suggestions = append(suggestions, command)
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

// Sets the path to be empty
func resetPath() string {
	context.Library = ""
	context.Topic = ""
	context.Key = ""
	return context.GetPath()
}

// Converts a path into the pop and push parts
// e.g. ../../actors/Priest has ../../ as the pop parts and actors/Priest as the push parts
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

// Gets the number of records that exist.  It is based on the depth of the path.
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

// Copies the current context so it can be reverted if an error occurs
func copyContext() {
	previousContext.Library = context.Library
	previousContext.Topic = context.Topic
	previousContext.Key = context.Key
	previousContext.Path = context.Path
}

// Revert the context to the values in previousContext
func revertContext() {
	context.Library = previousContext.Library
	context.Topic = previousContext.Topic
	context.Key = previousContext.Key
	context.Path = previousContext.Path
}

// Show the value for the record with the matcing library, topic, and key.
// If .showall is true, the json of the entire record will be shown.
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
				if settings.ShowAll {
					var jsonData []byte
					jsonData, err := json.Marshal(rec)
					if err != nil {
						log.Println(err)
					}
					fmt.Println(string(jsonData))
				} else {
					fmt.Println(rec.Value)
				}
			}
		}
	}
}

// escape quotes
func escape(value string) string {
	value = strconv.Quote(value)
	// a side effect of strconv.Quote is it wraps the entire string in quotes,
	// which we do not want.  So, remove them.
	value = value[1 : len(value)-1]
	return value
}

// converts value to int and stores it in settings.Padding for indicated index of 1 to 3
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

// set the context Library, Topic, and Key using the id
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

// changes the virtual directory based on parameters supplied
func changeDirectory(blocks []string) {
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
		up, dirs := cdParts(blocks[1])
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
			if mapper.Exists(context.Library, context.Topic, context.Key) {
				showValue(context.Library, context.Topic, context.Key)
				if settings.Hints {
					fmt.Printf("Hint: cmp to compare values for records with topic/key = %s/%s\n", context.Topic, context.Library)
				}
			} else {
				revertContext()
				fmt.Printf("%s not found\n", blocks[1])
			}
		} else {
			reportCount()
		}
	}
	setSuggestions()
}

// If the context is 3 levels deep, i.e. library/topic/key, displays value of all records with the same topic/key
func compareValues() {
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
			fmt.Printf("%4d %s\n", i+1, rec.ID)
			fmt.Printf("%s\n", strings.Repeat("-", context.TWidth))
			fmt.Println(rec.Value)
		}
		if settings.Hints {
			l := len(idMap.Map)
			fmt.Printf("\nHint: cd {number} to change directory, e.g. cd %d to change to %s\n", l, idMap.Map[l-1].ToPath())
		}
	} else {
		fmt.Println("You must be three levels deep to use this command.")
	}
}
func find(blocks []string) {
	conc.Map = make(map[string]concord.ConcordanceLine)
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
				sb.WriteString(pathDelimiter)
				sb.WriteString("%")
			case 2:
				sb.WriteString(context.Library)
				sb.WriteString(pathDelimiter)
				sb.WriteString(context.Topic)
				sb.WriteString(pathDelimiter)
				sb.WriteString("%")
			case 3:
				sb.WriteString(context.Library)
				sb.WriteString(pathDelimiter)
				sb.WriteString(context.Topic)
				sb.WriteString(pathDelimiter)
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

				conc.Line(fmt.Sprintf("%*s", settings.Padding.P2, rec.ID), line, value, settings.Width)
			}
			for i, res := range concord.SortedKeys(conc.Map, settings.Sort) {
				idMap.Add(i+1, conc.Map[res].ID)
				if len(strings.TrimSpace(conc.Map[res].ID)) > 50 {
					parts := strings.Split(conc.Map[res].ID, pathDelimiter)
					if len(parts) == 3 {
						fmt.Printf("%*d |%s%s%s%s\n", settings.Padding.P1, i+1, parts[0], pathDelimiter, parts[1], pathDelimiter)
						fmt.Printf("     | %*s | %-s%s%s\n", settings.Padding.P2, parts[2], conc.Map[res].Left, conc.Map[res].Key, conc.Map[res].Right)
					} else {
						// bad
						fmt.Printf("%*d | %*s | %-s%s%s\n", settings.Padding.P1, i+1, settings.Padding.P2, conc.Map[res].ID, conc.Map[res].Left, conc.Map[res].Key, conc.Map[res].Right)
					}
				} else {
					fmt.Printf("%*d | %*s | %-s%s%s\n", settings.Padding.P1, i+1, settings.Padding.P2, conc.Map[res].ID, conc.Map[res].Left, conc.Map[res].Key, conc.Map[res].Right)
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
			if settings.Hints {
				fmt.Printf("Hint: use cd {number} to change directory to a library/topic/key, e.g. cd %d\n", 2322)
			}
		}
	}
}
func showHelp() {
	for _, command := range commands {
		fmt.Printf("%s:\t%s\n", command.Text, command.Description)
	}
	fmt.Println("")
	fmt.Println("* indicates a command that is not available yet.")
	fmt.Println("Commands that start with a dot are for the settings.")
	fmt.Println("The settings affect the way the other commands behave.")
	fmt.Printf("As you type, suggestions will appear as a pop up.\nTab to set focus on the pop up.\nYou can scroll up or down using the arrow keys.\nUse the <Enter> key to select an item in the pop up.\n")
	fmt.Println("You can also use the up and down keys to scroll through previously issued commands.")
}

// show Topic-Key as LML Topic-Key. One use of the shell is when the user is creating templates and is trying to find
// a topic-Key.  This function will display the topic-key in the correct format
// for the liturgical markup language. Note that LML does not use libraries in the template.
// The libraries are prefixed to LML topic-keys when books/services are generated.
func showTKAsLML() {
	if context.Depth() == 3 {
		msg := fmt.Sprintf(`"%s%s%s"`, context.Topic, pathDelimiter, context.Key)
		fmt.Println(msg)
	} else {
		fmt.Println("You must be three levels deep to use this command")
	}
}

// Lists libraries if context.Depth() == 0, topics if == 1, keys if == 2, record value if == 3.
// If blocks[1] == "empty" it will list ids of records with an empty redirect and empty value column.
// If blocks[1] == "redirects" it will list ids of records with redirects.
func list(blocks []string) {
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
		fmt.Printf("%d empty records found\n", len(ids))
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
					idMap.Add(i+1, id.ToSQLId())
					fmt.Printf("%4d %s\n", i+1, library)
				}
				if settings.Hints {
					fmt.Printf("Hint: use cd {number} to change directory to that library, e.g. cd %d\n", len(libraries))
				}
			}
		case 1:
			{ // list topics for current library
				var like = context.Library
				if len(blocks) > 1 {
					like = like + pathDelimiter + blocks[1]
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
					idMap.Add(i+1, id.ToSQLId())
					fmt.Printf("%4d %s\n", i+1, topic)
				}
				if settings.Hints {
					fmt.Printf("Hint: use cd {number} to change directory to a library/topic, e.g. cd %d\n", len(topics))
				}
			}
		case 2:
			{ // list keys for current library/topic
				var like = context.Library + pathDelimiter + context.Topic
				if len(blocks) > 1 {
					like = like + pathDelimiter + blocks[1]
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
					idMap.Add(i+1, id.ToSQLId())
					fmt.Printf("%4d %s\n", i+1, key)
				}
				if settings.Hints {
					fmt.Printf("Hint: use cd {number} to change directory to a library/topic/key, e.g. cd %d\n", len(keys))
				}
			}
		case 3:
			{ // Display the value of the key
				showValue(context.Library, context.Topic, context.Key)
			}
		}
	}
}
// For the record whose id matches the context, sets the specified column
// which can be comment, redirect, or value.
func setRecord(blocks []string) {
	if context.Depth() < 3 {
		fmt.Println("You must be three levels deep to use the set command")
		return
	}
	rec, err := mapper.ReadByLTK(context.Library, context.Topic, context.Key)
	if err != nil {
		fmt.Println(err)
	} else {
		var prompt = fmt.Sprintf("set what, to what?\nExample 1: set value PRIEST\nExample 2: set redirect gr_gr_cog/template.titles/d.onSaturdayEvening\nExample 3: set comment kairos prayer\nTo set blank: set value '' or set value \"\"\n")
		value := ""
		if len(blocks) > 1 {
			if len(blocks) > 2 {
				value = escape(strings.Join(blocks[2:len(blocks)], " "))
				switch blocks[1] {
				case "comment":
					if value == "''" || value == "\\\"\\\"" {
						value = ""
					}
					rec.Comment = value
				case "redirect":
					if rec.Redirect == "''" || value == "\\\"\\\"" {
						value = ""
						rec.SetRedirect(value)
					} else {
						var id models.Id
						id.Parse(value)
						if ! mapper.Exists(id.Domain.ToNeo(), id.Topic, id.Key) {
							fmt.Printf("Does not exist in database: %s\n", value)
							return
						}
						if context.Library != id.Domain.ToNeo() {
							fmt.Println("You can only redirect to an ID in the same library")
							return
						}
						if context.Path == id.ToPath() {
							fmt.Println("You can not redirect a record to itself")
							return
						}
						rec.SetRedirect(value)
					}
				case "value":
					if value == "''" || value == "\\\"\\\"" {
						value = ""
					}
					rec.SetValue(value)
				default:
					fmt.Println(prompt)
				}
				err = mapper.Merge(rec)
				if err != nil {
					fmt.Println(err)
					return
				}
				switch blocks[1] {
				case "comment":
					fmt.Println(rec.Comment)
				case "redirect":
					if len(rec.Redirect) > 0 {
						idMap.Reset()
						idMap.Add(1, value)
					}
					fmt.Printf("Redirects to: 1 %s\n", rec.Redirect)
				case "value":
					fmt.Println(rec.Value)
				}
			} else {
				fmt.Println(prompt)
			}
		} else {
			fmt.Println(prompt)
		}
	}
}