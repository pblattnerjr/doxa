package atem2lml
/**
Temporary functions to convert atem files to lml.
TODO: remove this file and directory when Fr. S is using Doxa instead of ALWB
*/

import (
	"github.com/liturgiko/doxa/pkg/utils/ltfile"
)
var sections Sections

func Process(path string) error {
	files, err := ltfile.FileMatcher(path, "atem", nil)
	if err != nil {
		return err
	}
	for _, f := range files {
		sections.ProcessTemplate(f)
	}
	return nil
}
