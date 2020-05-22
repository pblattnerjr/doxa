// Package ltfile provides file utilities for processing liturgical texts
package ltfile

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

// FileExists returns true if the final segment of the path exists and is not a directory.
func FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
// DirExists returns true if the final segment of the path exists and is a directory
func DirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}
// CreateDir will create a directory if it does not exist.
// The error will be nil unless something is wrong with the path.
func CreateDir(path string) error {
	var err error = nil
	if ! DirExists(path) {
		err = os.Mkdir(path, os.ModePerm)
	}
	return err
}
// CreateDirs creates the specified directory and parents in path if they do not exist.
// The error will be nil unless something is wrong with the path.
func CreateDirs(path string) error {
	var err error = nil
	if ! DirExists(path) {
		err = os.MkdirAll(path, os.ModePerm)
	}
	return err
}
// WriteFile writes the supplied content using the filename.
func WriteFile(filename, content string) error {
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {return err}
	w := bufio.NewWriter(f)
	_, err = w.WriteString(content)
	w.Flush()
	return err
}

// HTMLTemplateToFile parses the html template and writes it to the filename provided.
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
// FileMatcher returns all files recursively in the dir that have the specified file extension and match the supplied regular expressions.
// The file extension should not have a leading dot.
// The expressions are regular expressions that will match the name of a file, but
// without the extension.  The extension will be automatically added to each expression.
// The expressions may be nil, in which case all file patterns will match.
func FileMatcher(dir, extension string, expressions []string) ([]string, error) {
	var result []string
	extensionPattern := "\\." + extension
	// precompile the expressions
	if expressions == nil {
		expressions = []string{".*"}
	}
	patterns := make([]*regexp.Regexp, len(expressions))
	for i, e := range expressions {
		p, err := regexp.Compile(e + extensionPattern)
		if err != nil {return result, err}
		patterns[i] = p
	}
	// now walk the files and apply the regular expressions
	err := filepath.Walk(ResolvePath(dir),
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
// Returns the directory within which the caller is executing
// This is just an example.  It must exist as a function in
// the desired package.
func executionDir() string {
	_, filename, _, _ := runtime.Caller(0)
	dir, _ := path.Split(filename)
	return dir
}
// ResolvePath checks the path for a leading tilde and returns the expanded path.
// The tilde is replaced with the user's home dir path.
// If no tilde is present, the path is returned unmodified.
func ResolvePath(path string) string {
	usr, _ := user.Current()
	dir := usr.HomeDir
	var newPath = ""
	if path == "~" {
		newPath = dir
	} else if strings.HasPrefix(path, "~/") {
		// Use strings.HasPrefix so we don't match paths like
		// "/something/~/something/"
		newPath = filepath.Join(dir, path[2:])
	} else {
		newPath = path
	}
	return newPath
}
// GetFileLines opens the file at the specified path and returns an array of its lines
func GetFileLines(path string) ([]string, error) {
	readFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()
	return fileTextLines, err
}