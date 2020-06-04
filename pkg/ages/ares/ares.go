/*
Provides utilities for parsing lines from ares files from the AGES Liturgical Workbench (ALWB) system. The acronymn `ares` means, `AGES resource.` The name of an ares file indicates its subject (called a topic) and its version (called a domain).
*/
package ares

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/liturgiko/doxa/pkg/utils/ltfile"
	"github.com/liturgiko/doxa/pkg/utils/ltstring"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)
var Logger *log.Logger
var (
	ErrFileMissingAres      = errors.New("ares: file missing .ares extension")
	ErrFileMissingTopic     = errors.New("ares: filename missing topic or domain")
	ErrLineMissingEqualSign = errors.New("ares: missing equal sign between key and value")
	ErrRedirectInvalid      = errors.New("ares: invalid redirect value")
	ErrValueMissingQuote    = errors.New("ares: value missing initial or final quote")
)

func NewErrLineMissingEqualSign(line int) error {
	return errors.New(fmt.Sprintf("line %d ares: missing equal sign between key and value", line))
}
func NewErrValueMissingQuote(line int) error {
	return errors.New(fmt.Sprintf("line %d ares: value missing initial or final quote", line))
}

type LineParts struct {
	Language       string
	Country        string
	Realm          string
	Topic          string
	Key            string
	Value          string
	Redirect       string
	Comment        string
	IsAresId       bool
	IsBlank        bool
	IsCommentedOut bool
	IsRedirect     bool
	HasComment     bool
	HasValue       bool
	LineNbr        int
}

type FilenameParts struct {
	Topic    string
	Language string
	Country  string
	Realm    string
}

// Parse the line into the fields of a LineParts struct.
// We include the language, country, realm, and topic
// so they are available to downstream channels
func ParseLine(filenameParts FilenameParts, line string) (LineParts, error) {
	var result LineParts
	result.Language = filenameParts.Language
	result.Country = filenameParts.Country
	result.Realm = filenameParts.Realm
	result.Topic = filenameParts.Topic
	result.IsAresId = strings.HasPrefix(line, "A_")
	result.IsBlank = len(strings.TrimSpace(line)) < 1
	result.IsCommentedOut = strings.HasPrefix(line, "//")
	if result.IsBlank || result.IsAresId || result.IsCommentedOut {
		// fall thru
	} else {
		parts := strings.Split(line, "=")
		if len(parts) > 1 {
			result.Key = strings.TrimSpace(parts[0])
			result.Value = strings.TrimSpace(parts[1])
			startComment := strings.Index(result.Value, "//")
			lastQuote := strings.LastIndex(result.Value, "\"")
			if startComment > -1 && lastQuote < startComment { // The string "//" is valid inside a definition
				result.HasComment = true
				valueParts := strings.Split(result.Value, "//")
				result.Value = strings.TrimSpace(valueParts[0])
				result.Comment = strings.TrimSpace(valueParts[1])
			}
			result.IsRedirect = !strings.HasPrefix(result.Value, "\"")
			if result.IsRedirect {
				result.Redirect, _ = ToRedirectId(result.Value)
				result.Value = ""
			} else {
				if strings.HasPrefix(result.Value, "\"") && strings.HasSuffix(result.Value, "\"") {
					result.Value, _ = strconv.Unquote(result.Value)
					result.HasValue = true
				} else {
					return result, ErrValueMissingQuote
				}
			}
		} else {
			return result, ErrLineMissingEqualSign
		}
	}
	return result, nil
}

// Returns from the filename the topic, ISO country code, ISO language code, and realm.
func ParseAresFileName(f string) (FilenameParts, error) {
	var result FilenameParts
	_, filename := filepath.Split(f)
	if strings.HasSuffix(filename, ".ares") {
		parts := strings.Split(filename[:len(filename)-5], "_")
		if len(parts) == 4 {
			result.Topic = parts[0]
			result.Language = strings.ToLower(parts[1])
			result.Country = strings.ToLower(parts[2])
			result.Realm = strings.ToLower(parts[3])
		} else {
			return result, ErrFileMissingTopic
		}
	} else {
		return result, ErrFileMissingAres
	}
	return result, nil
}

// Returns the library name (Language_Country_Realm) from an ares filename
func LibraryName(f string) (string, error) {
	var result string
	fparts, err := ParseAresFileName(f)
	if err != nil {
		return result, err
	} else {
		result = fparts.Language + "_" + fparts.Country + "_" + fparts.Realm
		return result, nil
	}
}

// Converts value from this form: properties_en_UK_lash.media.key
// To this form: en_us_lash~properties~media.key
func ToRedirectId(value string) (string, error) {
	var lang, country, realm, topic, key string
	domainParts := strings.Split(strings.TrimSpace(value), "_")
	if len(domainParts) == 4 {
		topic = domainParts[0]
		lang = domainParts[1]
		country = domainParts[2]
		keyParts := strings.Split(domainParts[3], ".")
		realm = keyParts[0]
		key = domainParts[3][len(realm)+1:]
		return ltstring.ToId(lang, country, realm, topic, key), nil
	} else {
		return value, ErrRedirectInvalid
	}
}

// CleanAres cleans Ares files by finding and fixing the following problems:
//
// When finds a value that starts with quote but does not end with one,
// attempts to joint the next line to it, if the next line appears to be the
// broken part (due to a line break)
//
// When finds a duplicate key, compares the values and keeps the key that has a value,
// throwing away the key that does not.
// If both have values, reports the problem in the log.
// If only one has a value, deletes the key that does not have a value.
// Implements F.2019.005

const emptyString = "\"\""

var lineCnt int
var fileIn *os.File
var fileOut *os.File

var InBase string	// base path of all input files
var OutBase string	// base path of all output files


// The following create global lists of definitions and references to other keys found in an
// entire run of any program cleaning a tree of ares files.

type aresReference struct {
	file string
	line int
}
type libraryReferences map[string] []aresReference       // count the number of times this reference is found
var  AllReferences  = make(map[string]libraryReferences) // all references for all libraries

type aresReferences map[string] string
var  AllDefinitions  = make(map[string]aresReferences) // all references for all libraries

func saveDefinition(k string, v string, c string, noComment bool,
	definitions map[string]string, commentsForKey map[string]string) bool {

	var found bool
	var msg string
	var oldValue string

	oldValue, found = definitions[k]

	if found == false {
		// Case 1: no previous definition, value is "" - save value
		// Case 2: no previous definition, value is NOT "" - save value
		definitions[k] = v // new value
		if noComment && Logger != nil {
			Logger.Println(fmt.Sprintf("%s %s", k, c))
		} else {
			commentsForKey[k] += c
		}
	} else if oldValue == emptyString && v == emptyString {
		// Case 3: saved definition and current definition are both "" - discard this definition
		msg = fmt.Sprintf("// line %d duplicate empty definition for key %s - discarded\n", lineCnt, k)
		if noComment && Logger != nil  {
			Logger.Println(fmt.Sprintf("%s %s", k, c + msg))
		} else {
			commentsForKey[k] += c + msg
		}
	} else if oldValue == emptyString && v != emptyString {
		// Case 4: replacing empty definition with real definition
		definitions[k] = v // substitute real value for placeholder
		msg = fmt.Sprintf("// line %d empty definition for key %s replaced with %s\n", lineCnt, k, v)
		if noComment  && Logger != nil {
			Logger.Println(fmt.Sprintf("%s %s", k, c + msg))
		} else {
			commentsForKey[k] += c + msg
		}
	} else if oldValue != emptyString && v == emptyString {
		// Case 5: old definitions is non-empty but new definition is "" - discard new definition
		msg = fmt.Sprintf("// line %d Invalid replacement of value %s for key %s with empty string\n", lineCnt, oldValue, k)
		if noComment  && Logger != nil {
			Logger.Println(fmt.Sprintf("%s %s", k, c + msg))
		} else {
			commentsForKey[k] += c + msg
		}
	} else if oldValue != emptyString && v != emptyString {
		if v == oldValue {
			// Case 6: both old definitions and new definition are non-empty and the same - discard new definition
			msg = fmt.Sprintf("// line %d duplicate value %s for key %s - discarded\n", lineCnt, v, k)
		} else {
			// Case 7: both old definitions and new definition are non-empty - use new definition
			msg = fmt.Sprintf("// line %d Substituting value %s for key %s old value was %s\n", lineCnt, v, k, oldValue)
			definitions[k] = v // updating value for key
		}
		if noComment && Logger != nil  {
			Logger.Println(fmt.Sprintf("%s %s", k, c + msg))
		} else {
			commentsForKey[k] += c + msg
		}
	}
	return found
}

func alreadyDefined(k string, definitions map[string]string) bool {
	var found bool
	_, found = definitions[k]
	return found
}

func hasComments(k string, commentsForKey map[string]string) bool {
	var found bool
	_, found = commentsForKey[k]
	return found
}
func CleanAresFiles(dirIn, dirOut string, noComment bool, logger *log.Logger) error {
	Logger = logger
	dirIn = ltfile.ResolvePath(dirIn)
	dirOut = ltfile.ResolvePath(dirOut)

	// filter for only files that have a domain, e.g. gr_gr_cog
	var expressions = []string{
		".*_.*_.*",
	}
	files, err := ltfile.FileMatcher(dirIn, "ares", expressions)
	if err != nil {
		return errors.New(fmt.Sprintf("no ares files in %s\n", dirIn))
	}

	for _, f := range files {
		err = CleanAres(f, strings.Replace(f,dirIn, dirOut,1), noComment)
		if err != nil {
			Logger.Println(err)
		}
	}
	return nil
}

func CleanAres(in, out string, noComment bool) error {
	var (
		err            error
		parts          []string
		library        string
		line           string
		nextLine       string
		aresKey        string
		aresDef        string
		keys           []string
		comments       string
		definitions    map[string]string
		commentsForKey map[string]string
	)

	definitions = make(map[string]string)
	commentsForKey = make(map[string]string)
	comments = ""

	fileIn, err := os.Open(in)
	if err != nil {
		Logger.Fatal(err)
	}
	defer fileIn.Close()
	scanner := bufio.NewScanner(fileIn)

	theDir := filepath.Dir(out)
	err = ltfile.CreateDirs(theDir)
	if err != nil {
		Logger.Fatal(err)
	}
	fileOut, err := os.Create(out)
	if err != nil {
		Logger.Fatal(err)
	}
	defer fileOut.Close()
	w := bufio.NewWriter(fileOut)

	library, err = LibraryName(in)
	if err != nil {
		if Logger != nil {
			Logger.Printf("%s %v", in, ErrFileMissingTopic)
		} else {
			Logger.Printf("%s %v", in, ErrFileMissingTopic)
		}
		return nil
	}

	var blockComment bool

	//
	// Phase 1 - scan the file creating a slice of keys to be written and
	// a map of definitions index by the keys
	//
	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
		line = strings.TrimSpace(line)
		lineCnt++

		// Append all blank or comment lines to the 'comments' variable
		// to be associated with the next key
		if len(line) == 0 || strings.HasPrefix(line, "//") {
			comments = comments + line + "\n"
			continue
		}
		if strings.HasPrefix(line, "/*") || blockComment {
			comments = comments + line + "\n"
			blockComment = true
			continue
		}
		if strings.HasPrefix(line, "*/") || strings.HasSuffix(line, "*/") {
			comments = comments + line + "\n"
			blockComment = false
			continue
		}

		if !strings.Contains(line, "=") { // no "=" in the line, just treat it like a comment
			comments = comments + line + "\n"
			continue
		} else { // this is an assignment of some sort
			parts = strings.Split(line, "=")
			aresKey = strings.TrimSpace(parts[0])
			aresDef = strings.TrimSpace(parts[1])
			aresDef = strings.ReplaceAll(aresDef, "\\\"", "%%") // guard escaped quotes
			if len(parts) > 2 {                                 // a few lines have an '=' in a trailing comment
				aresDef = aresDef + " = " + strings.TrimSpace(parts[2])
			}

			if strings.HasPrefix(aresDef, "\"") { // right side starts with quote
				// Special multi-line case with embedded newline
				// if aresDef has no close quote
				if aresDef == "\"" || !strings.HasSuffix(aresDef, "\"") && strings.Count(aresDef, "\"") == 1 {
					for {
						scanner.Scan()
						nextLine = strings.TrimSpace(scanner.Text())
						nextLine = strings.ReplaceAll(nextLine, "\\\"", "%%") // guard escaped quotes
						lineCnt++
						if aresDef != "\"" && nextLine != "\"" {
							aresDef = aresDef + " " + nextLine
						} else {
							aresDef = aresDef + nextLine // no need for extra space at start or end
						}

						if strings.HasSuffix(nextLine, "\"") { // This is the last line of the definition.
							break
						}
					}
				}
				if _,ok := AllDefinitions[library][aresKey]; !ok { // create a new map of definitions
					AllDefinitions[library] = make(map[string] string)
				}
				AllDefinitions[library][aresKey] = aresDef

			} else {                                              // without quotes, this is a reference, not a definition.  Store it for checking later
				if _,ok := AllReferences[library][aresKey]; !ok { // of there is no entry for the key reference
					refs := make(libraryReferences)
					AllReferences[library] = refs
				}
				var aReference aresReference
				aReference.file = in[len(InBase)+1:]
				aReference.line = lineCnt
				AllReferences[library][aresKey] = append(AllReferences[library][aresKey],aReference)
			}

			if !alreadyDefined(aresKey, definitions) { // add only the first time the key is encountered
				keys = append(keys, aresKey)
			}

			aresDef = strings.ReplaceAll(aresDef, "%%", "\\\"") // restore escaped quotes

			saveDefinition(aresKey, aresDef, comments, noComment, definitions, commentsForKey)
			comments = ""
		}
	}

	//
	// phase 2 - using the slice of keys, write out the keys and their definitions
	//
	lineCnt = 0 // reset to zero

	for _, k := range keys {
		if hasComments(k, commentsForKey) { // are there comments for the key to write out first?
			w.WriteString(commentsForKey[k])
		}
		w.WriteString(k + " = " + definitions[k] + "\n")
	}
	if len(comments) > 0 {
		w.WriteString(comments) // These lines were found at the end of the file (have no key).
	}
	w.Flush()
	return err
}

// GetAresErrors returns an array of errors found
// in the specified ares output file.  The errors will
// be one of the error types defined in this package.
func GetAresErrors(out string) *[]error {
	var result []error
	file, err := os.Open(out)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var fileNameParts FilenameParts
	fileNameParts, err = ParseAresFileName(file.Name())
	fname := file.Name()
	fname = fname[len(OutBase)+2:]
	if err != nil {
		result = append(result, errors.New(fmt.Sprintf("%s: %v", fname, err)))
	}
	var lineCnt int
	inCommentBlock := false

	seenKey := map[string]bool{}

	for scanner.Scan() {
		lineCnt = lineCnt + 1
		line := strings.TrimSpace(scanner.Text())
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "/*") {
			inCommentBlock = true
			continue
		}
		if strings.HasPrefix(line, "*/") || strings.HasSuffix(line, "*/") {
			inCommentBlock = false
			continue
		}
		if !inCommentBlock {
			lineParts, err := ParseLine(fileNameParts, line)
			lineParts.LineNbr = lineCnt
			if len(lineParts.Key) > 0 && seenKey[lineParts.Key] {
				result = append(result, errors.New(fmt.Sprintf("%s Line: %d: duplicate key %s", fname, lineCnt, lineParts.Key)))
			} else {
				seenKey[lineParts.Key] = true
			}
			if err != nil {
				result = append(result, errors.New(fmt.Sprintf("%s line: %d: %s", fname, lineCnt, err)))
			}
		}
	}
	return &result
}
