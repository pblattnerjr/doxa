package atem2lml

/**
Temporary functions to convert atem files to lml.
TODO: remove this file and directory when Fr. S is using Doxa instead of ALWB
*/

import (
	"fmt"
	"github.com/google/martian/log"
	"github.com/liturgiko/doxa/pkg/utils/ltfile"
	"os"
	"strings"
)

// Process converts all .atem files (templates) found in dirIn to .lml files, written to dirOut.
// .atem files are AGES templates.  .lml files are Doxa liturgical markup language templates.
// A non-nil error is returned if the path dirIn does not exist.
// Precondition to calling this function: the liturgical database must
// have already been loaded from the .ares files.
func Process(dirIn, dirOut string) error {
	// TOL
	// Since it appears that all inserts are fully qualified, maybe what I need to do is:
	// 1. Flatten templates so that each section is in a subdirectory
	//    a. write out the template as a .lml
	//    b. when encounter a section,
	//   	1. push current file writer on stack
	//  	2. open a new file writer
	//    c. when encounter end-section
	//      1. close the file writer
	//      2. set file writer to stack pop
	// 2. Convert insert paths to use / instead of .
	// Regarding ares keys,
	// Perhaps you do a read of the db for library = gr_gr_cog and find the key.
	// e.g.
	// select topic || "/" || key from ltx where library = "gr_gr_cog" and key = "hoMA.LaudsSunTheotokion.text";
	// If multiple hits occur, examine the map of imports for the template and find
	// one that matches.  Note that there might be fully qualified keys, so watch out for them.

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
