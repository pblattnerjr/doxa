// Package config provides a means to initialize a config file for the doxago cli
package config

import (
	"fmt"
	"github.com/liturgiko/doxa/pkg/utils/ltfile"
	"github.com/liturgiko/doxa/pkg/utils/repos"
	"log"
	"os"
	"path/filepath"
	"text/template"
)
const (
	AtemDir = "atem"
	DataDir = "data"
	HTTPDir = "http"
	LogDir = "logs"
	ReposDir = "repos"
	Site = "site"
	SQL = "sql"
	SQLDbName = "liturgical.db"
	StaticDir = "liml"
	SysDir = "sys"
	TemplatesDir = "templates"
)
type Paths struct {
	AtemPath string // AGES atem files
	DbPath string
	HomePath string
	HTTPPath string
	LogPath string
	ReposPath string
	SitePath string
	SysPath string // AGES ares system files
	TemplatesPath string
}
func NewPaths(home string, siteUrl string) *Paths {
	var paths = new(Paths)
	paths.HomePath = home
	paths.DbPath = filepath.Join(paths.HomePath, DataDir, SQL, SQLDbName)
	paths.HTTPPath = filepath.Join(paths.HomePath, HTTPDir)
	paths.LogPath = filepath.Join(paths.HomePath, LogDir)
	paths.ReposPath = filepath.Join(paths.HomePath, ReposDir)
	sitePath, _ := repos.DirPath(filepath.Join(paths.ReposPath, Site),siteUrl)
	paths.SitePath = sitePath
	paths.TemplatesPath = filepath.Join(paths.HomePath, TemplatesDir)
	paths.AtemPath = filepath.Join(paths.ReposPath, AtemDir)
	paths.SysPath = filepath.Join(paths.ReposPath, SysDir)
	return paths
}
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
test.github.repos.ares:
- https://github.com/liturgiko/testrepo1.git
- https://github.com/liturgiko/testrepo2.git

# Github repositories to be processed
# For the clone, each one is processed concurrently.
# So list the largest first in order to reduce processing time.
github.repos.ares:
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
- https://github.com/AGES-Initiatives/ages-alwb-library-en-us-duvall.git
- https://github.com/AGES-Initiatives/ages-alwb-library-en-us-houpos.git
- https://github.com/AGES-Initiatives/ages-alwb-library-en-us-repass.git
- https://github.com/AGES-Initiatives/ages-alwb-library-en-us-unknown.git
- https://github.com/AGES-Initiatives/ages-alwb-library-lash.git
- https://github.com/AGES-Initiatives/ages-alwb-scripture.git

# OSLW repo (optional)
github.repos.oslw: https://github.com/OCMC-Translation-Projects/service.book.ke.oak.git

# template repo (optional). Used for converting ares templates to lml templates
github.repos.atem: https://github.com/AGES-Initiatives/ages-alwb-templates.git

# system repo (optional). Used for converting ares templates to lml templates
github.repos.sys: https://github.com/AGES-Initiatives/ages-alwb-system.git

# glory is the index of which language to use to give glory.
# The index starts with 0 (for Greek)
glory: 0
glory.languages:
- Δόξα Πατρὶ καὶ Υἱῷ καὶ Ἁγίῳ Πνεύματι, νῦν καὶ ἀεὶ καὶ εἰς τοὺς αἰῶνας τῶν αἰώνων!
- Gloria al Padre, y al Hijo y al Espíritu Santo, perpetuamente, ahora y siempre y por los siglos de los siglos!
- Utukufu kwa Baba na Mwana na Roho Mtakatifu, Daima, sasa na siku zote, hata milele na milele!
- Glory to the Father and the Son and the Holy Spirit, both now and forever and unto the ages of ages!)

--- # document end

`
