package atem2lml

/**
Temporary functions to convert atem files to lml.
TODO: remove this file and directory when Fr. S is using Doxa instead of ALWB
*/

import (
	"fmt"
	"github.com/google/martian/log"
	"github.com/liturgiko/doxa/pkg/ages/tags"
	"github.com/liturgiko/doxa/pkg/utils/ltfile"
	"os"
	"strings"
)

var ITags tags.Tag // inline tags
//var BTags tags.Tag // block tags
//var RoleTags tags.Tag
//var RuleTags tags.Tag

// Process converts all .atem files (templates) found in dirIn to .lml files, written to dirOut.
// .atem files are AGES templates.  .lml files are Doxa liturgical markup language templates.
// A non-nil error is returned if the path dirIn does not exist.
// Precondition to calling this function: the liturgical database must
// have already been loaded from the .ares files.
func Process(sysDir, dirIn, dirOut string) error {
	// make sure dirOut ends with path delimiter
	pathSeparator := string(os.PathSeparator)
	if !strings.HasSuffix(dirOut, pathSeparator) {
		dirOut = dirOut + pathSeparator
	}
	// empty the lml templates directory and create a new one
	err := os.RemoveAll(dirOut)
	if err != nil {
		return err
	}
	ltfile.CreateDirs(dirOut)

	// load the tag ares files
	ITags.Load(sysDir, "iTags")
	//BTags.Load(sysDir, "bTags")
	//RoleTags.Load(sysDir, "roles")
	//RuleTags.Load(sysDir, "rules")

	files, err := ltfile.FileMatcher(dirIn, "atem", nil)
	if err != nil {
		return err
	}
	library := "ages/"
	// Create an index of all templates and sections so ParseTemplate can get the path to any sections inserted.
	tTree, sTree, err := Index(dirIn,library)
	// Process each AGES template and write out the corresponding lml files.
	// Each section will be written out as a separate file.
	for _, f := range files {
		parts := strings.Split(f, "a-templates/")
		if len(parts) == 2 {
			id := library + parts[1][0:len(parts[1])-5]
			if err := ParseTemplate(f, dirOut, id, tTree, sTree); err != nil {
				fmt.Println(err)
			}
		} else {
			log.Errorf("Should be using a-templates in path %s\n", f)
		}
	}
	return nil
}
