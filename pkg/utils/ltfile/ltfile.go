// file utilities for processing liturgical texts
package ltfile

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

// FileExists returns true if final segment of the path exists and is not a directory.
// path is the path to the file and the filename as the final part of the path.
func FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
// DirExists returns true of the final segment of the path exists and is a directory
func DirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
// CreateDir will create a directory if it does not exist.
// If it does exist, error is nil.
func CreateDir(path string) error {
	var err error = nil
	if ! DirExists(path) {
		err = os.Mkdir(path, os.ModePerm)
	}
	return err
}
// CreateDirs will create directory and parents in path if they do not exist.
// If they does exist, error is nil.
func CreateDirs(path string) error {
	var err error = nil
	if ! DirExists(path) {
		err = os.MkdirAll(path, os.ModePerm)
	}
	return err
}
func WriteFile(filename, content string) error {
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {return err}
	w := bufio.NewWriter(f)
	_, err = w.WriteString(content)
	w.Flush()
	return err
}
// Parses the html template and writes it to the filename provided.
// name: the name of the template
// html: the template
// filename: the name of the html file including its path
// data: the data to be filled into the html (may be nil)
func HTMLTemplateToFile(name, html, filename string, data interface{}) error {
	t := template.Must(template.New(name).Parse(html))
	f, err := os.Create(filename)
	if err != nil {
		log.Println(fmt.Sprintf("create file %s: %s", filename, err))
		return err
	}
	err = t.Execute(f, data)
	if err != nil {
		log.Println(fmt.Sprintf("create template: %s", err))
		return err
	}
	f.Close()
	return nil
}
// Returns all files recursively in the dir that have the specified file extension
// without a dot, and match any one of the supplied expressions.
// The expressions are regular expressions that will match the name of a file, but
// without the extension.  The extension will be automatically added to each expression.
func FileMatcher(dir, extension string, expressions []string) ([]string, error) {
	var result []string
	extensionPattern := "\\." + extension
	// precompile the expressions
	patterns := make([]*regexp.Regexp, len(expressions))
	for i, e := range expressions {
		p, err := regexp.Compile(e + extensionPattern)
		if err != nil {return result, err}
		patterns[i] = p
	}
	// now walk the files and apply the regular expressions
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			for _, p := range patterns {
				if p.MatchString(info.Name()) {
					result = append(result, path)
				}
			}
			return nil
		})
	return result, err
}

func getWalkFunc(patterns []*regexp.Regexp) filepath.WalkFunc {
	return func(path string, fileInfo os.FileInfo, err error) error {
		// ...do something with service...
		return nil
	}
}

