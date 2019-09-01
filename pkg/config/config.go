// Package config provides a means to initialize a config file for the doxago cli
package config

import (
	"fmt"
	"github.com/liturgiko/doxa/pkg/utils/ltfile"
	"log"
	"os"
	"path/filepath"
	"text/template"
)
const (
	DataDir = "data"
	SQL = "sql"
	HTTPDir = "http"
	StaticDir = "liml"
	TemplatesDir = "templates"
	LogDir = "logs"
)
// Initializes the config file and directories
// for the doxago cli. If these already exist,
// they will not be overwritten.
func Initdoxa(home string) (string, error) {
	reposHome := filepath.Join(home, "repos")

	// create the various dirs
	if err := ltfile.CreateDir(home); err != nil {return "", err}
	if err := ltfile.CreateDirs(reposHome); err != nil {return "", err}
	if err := ltfile.CreateDirs(filepath.Join(home,DataDir)); err != nil {return "", err}
	if err := ltfile.CreateDirs(filepath.Join(home,DataDir, SQL)); err != nil {return "", err}
	if err := ltfile.CreateDirs(filepath.Join(home,HTTPDir)); err != nil {return "", err}
	if err := ltfile.CreateDirs(filepath.Join(home,LogDir)); err != nil {return "", err}
	if err := ltfile.CreateDirs(filepath.Join(home,TemplatesDir)); err != nil {return "", err}

	configPath := filepath.Join(home, ".doxago.yaml")

	if ! ltfile.FileExists(configPath) {
		// generate the config file
		t := template.Must(template.New("config").Parse(doxagoTmpl))
		f, err := os.Create(configPath)
		if err != nil {
			log.Println("create file: ", err)
			return "", err
		}
		if err := t.Execute(f,reposHome); err != nil {return "", err}
	}
	return fmt.Sprintf("doxa home directory is: %s\nConfig file is: %s", home, configPath), nil
}

const doxagoTmpl = `
--- # document start
# Settings for the doxago cli

# Directories
dir.repos: {{.}}

# Ports
port.http.doxa: 8080

# Generation settings
generate.domains:
- gr_gr_cog
- en_us_goarch
generate.file.pattern: eu*
generate.output.types:
- epub
- html
- pdf
# generate.pdf.lib values are: go, latex
# if you set it to latex, then latex and xelatex must be installed
# as well as the oslw libraries.
generate.pdf.lib: go

# Test Github repositories to be processed
test.github.repos:
- https://github.com/liturgiko/testrepo1.git
- https://github.com/liturgiko/testrepo2.git

# Github repositories to be processed
# For the clone, each one is processed concurrently.
# So list the largest first in order to reduce processing time.
github.repos:
- https://github.com/AGES-Initiatives/ages-alwb-library-gr-gr-cog.git
- https://github.com/AGES-Initiatives/ages-alwb-library-en-us-dedes.git
- https://github.com/AGES-Initiatives/ages-alwb-library-en-us-goa.git
- https://github.com/AGES-Initiatives/ages-alwb-library-en-us-holycross.git
- https://github.com/AGES-Initiatives/ages-alwb-library-en-us-oca.git
- https://github.com/AGES-Initiatives/ages-alwb-library-ancillary.git
- https://github.com/AGES-Initiatives/ages-alwb-library-client-enpublic.git
- https://github.com/AGES-Initiatives/ages-alwb-library-en-fr-omvol.git
- https://github.com/AGES-Initiatives/ages-alwb-library-en-uk-ware.git
- https://github.com/AGES-Initiatives/ages-alwb-library-en-us-acook.git
- https://github.com/AGES-Initiatives/ages-alwb-library-en-us-andronache.git
- https://github.com/AGES-Initiatives/ages-alwb-library-en-us-barrett.git
- https://github.com/AGES-Initiatives/ages-alwb-library-en-us-boyer.git
- https://github.com/AGES-Initiatives/ages-alwb-library-en-us-constantinides.git
- https://github.com/AGES-Initiatives/ages-alwb-library-en-us-duvall.git
- https://github.com/AGES-Initiatives/ages-alwb-library-en-us-houpos.git
- https://github.com/AGES-Initiatives/ages-alwb-library-en-us-repass.git
- https://github.com/AGES-Initiatives/ages-alwb-library-en-us-unknown.git
- https://github.com/AGES-Initiatives/ages-alwb-library-lash.git
- https://github.com/AGES-Initiatives/ages-alwb-scripture.git

--- # document end

`
