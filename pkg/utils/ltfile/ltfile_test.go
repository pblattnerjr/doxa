package ltfile

import (
	"fmt"
	"testing"
)

func TestFileMatcher(t *testing.T) {
	dir := "/Users/mac002/git/ages/atem/ages-alwb-templates/net.ages.liturgical.workbench.templates/a-templates/Books"
	ext := "atem"
	patterns := []string{
		"bk.liturgy.*",
	}
	files, err := FileMatcher(dir, ext, patterns)
	if err != nil {
		t.Error(err.Error())
	}
	if len(files) == 0 {
		t.Error("No matching files")
	}
	for _, file := range files {
		fmt.Println(file)
	}
}
