/*
Provides utilities for parsing lines from ares files from the AGES Liturgical Workbench (ALWB) system. The acronymn `ares` means, `AGES resource.` The name of an ares file indicates its subject (called a topic) and its version (called a domain).
*/
package ares

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/liturgiko/doxa/pkg/utils/ltstring"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

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
			if strings.Index(result.Value, "//") > -1 {
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

func saveDefinition(k string, v string, c string,
	definitions map[string]string, commentsForKey map[string]string, keys *[]string, fname string) bool {

	var found bool
	var msg string
	var oldValue string

	oldValue, found = definitions[k]

	if found == false {
		// Case 1: no previous definition, value is "" - save value
		// Case 2: no previous definition, value is NOT "" - save value
		definitions[k] = v // new value
		commentsForKey[k] += c
	} else if oldValue == emptyString && v == emptyString {
		// Case 3: saved definition and current definition are both "" - discard this definition
		msg = fmt.Sprintf("// line %d duplicate empty definition for key %s - discarded\n", lineCnt, k)
		commentsForKey[k] += c + msg
	} else if oldValue == emptyString && v != emptyString {
		// Case 4: replacing empty definition with real definition
		definitions[k] = v // substitute real value for placeholder
		msg = fmt.Sprintf("// line %d empty definition for key %s replaced with %s\n", lineCnt, k, v)
		commentsForKey[k] += c + msg
	} else if oldValue != emptyString && v == emptyString {
		// Case 5: old definitions is non-empty but new definition is "" - discard new definition
		msg = fmt.Sprintf("// line %d Invalid replacement of value %s for key %s with empty string\n", lineCnt, oldValue, k)
		commentsForKey[k] += c + msg
	} else if oldValue != emptyString && v != emptyString {
		if v == oldValue {
			// Case 6: both old definitions and new definition are non-empty and the same - discard new definition
			msg = fmt.Sprintf("// line %d duplicate value %s for key %s - discarded\n", lineCnt, v, k)
		} else {
			// Case 7: both old definitions and new definition are non-empty - use new definition
			msg = fmt.Sprintf("// line %d Substituting value %s for key %s old value was %s\n", lineCnt, v, k, oldValue)
			definitions[k] = v // updating value for key
		}
		commentsForKey[k] += c + msg
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

func CleanAres(in, out string) error {
	var (
		err            error
		parts          []string
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
		log.Fatal(err)
	}
	defer fileIn.Close()
	scanner := bufio.NewScanner(fileIn)

	theDir := filepath.Dir(out)
	err = os.MkdirAll(theDir,os.ModeDir)
	if err != nil {
		log.Fatal(err)
	}
	fileOut, err := os.Create(out)
	if err != nil {
		log.Fatal(err)
	}
	defer fileOut.Close()
	w := bufio.NewWriter(fileOut)

	var blockComment bool

	//
	// Phase 1 - scan the file creating a slice of keys to be written and
	// a map of definitions index by the keys
	//
	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
		line = strings.TrimSpace(line) // ??? why
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

		if !strings.Contains(line,"=") { // no "=" in the line, just treat it like a comment
			comments = comments + line + "\n"
			continue
		} else { // this is an assignment of some sort
			parts = strings.Split(line, "=")
			aresKey = strings.TrimSpace(parts[0])
			aresDef = strings.TrimSpace(parts[1])
			if len(parts) > 2 { // a few lines have an '=' in a trailing comment
				aresDef = aresDef + " = " + strings.TrimSpace(parts[2])
			}

			if strings.HasPrefix(aresDef, "\"") { // right side starts with quote
				// Special multi-line case with embedded newline
				// if aresDef has no close quote
				if !strings.HasSuffix(aresDef, "\"") && strings.Count(aresDef,"\"") == 1 {
					for {
						scanner.Scan()
						nextLine = strings.TrimSpace(scanner.Text())
						lineCnt++
						aresDef = aresDef + " " + nextLine
						if strings.Contains(nextLine, "\"") { // This is the next line of the definition.
							break
						}
					}
				}
			}

			if !alreadyDefined(aresKey, definitions) { // add only the first time the key is encountered
				keys = append(keys, aresKey)
			}

			saveDefinition(aresKey, aresDef, comments, definitions, commentsForKey, &keys, fileIn.Name())
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
// in the specified ares file.  The errors will
// be one of the error types defined in this package.
func GetAresErrors(in string) *[]error {
	var result []error
	file, err := os.Open(in)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var fileNameParts FilenameParts
	fileNameParts, err = ParseAresFileName(file.Name())
	if err != nil {
		result = append(result, err)
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
				result = append(result, errors.New(fmt.Sprintf("%s %d: duplicate key %s", file.Name(), lineCnt, lineParts.Key)))
			} else {
				seenKey[lineParts.Key] = true
			}
			if err != nil {
				result = append(result, errors.New(fmt.Sprintf("%s %d: %s", file.Name(), lineCnt, err)))
			}
		}
	}
	return &result
}
