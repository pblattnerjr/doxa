package ares

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"testing"
)
var fnp FilenameParts

func init() {
fnp, _ = ParseAresFileName("actors_gr_gr_cog.ares")
}

// Test with key and value
func TestLineHasKV(t *testing.T) {
	l := "x = \"y\""
	lineParts, err := ParseLine(fnp,l)
	if err != nil {
		t.Error("expected no error, got", err.Error())
	}
	if lineParts.HasComment {
		t.Error("Expected not hasComment, got", lineParts.HasComment)
	}
	if !lineParts.HasValue {
		t.Error("Expected hasValue, got", lineParts.HasValue)
	}
	if lineParts.IsAresId {
		t.Error("Expected not isAresId, got", lineParts.IsAresId)
	}
	if lineParts.IsBlank {
		t.Error("Expected ! isBlank, got", lineParts.IsBlank)
	}
	if lineParts.IsCommentedOut {
		t.Error("Expected ! isCommentedOut, got", lineParts.IsCommentedOut)
	}
	if lineParts.IsRedirect {
		t.Error("Expected ! isRedirect, got", lineParts.IsRedirect)
	}
	if lineParts.Key != "x" {
		t.Error("Expected key = x, got", lineParts.Key)
	}
	if lineParts.Value != "y" {
		t.Error("Expected value = y, got", lineParts.Value)
	}
	if len(lineParts.Comment) > 0 {
		t.Error("Expected len(comment) == 0, got", lineParts.Comment)
	}
}

// Test with redirect
func TestLineHasRedirect(t *testing.T) {
	l := "x = actors_gr_GR_cog.Priest"
	lineParts, err := ParseLine(fnp, l)
	if err != nil {
		t.Error("expected no error, got", err.Error())
	}
	if lineParts.HasComment {
		t.Error("Expected not hasComment, got", lineParts.HasComment)
	}
	if lineParts.HasValue {
		t.Error("Expected not hasValue, got", lineParts.HasValue)
	}
	if lineParts.IsAresId {
		t.Error("Expected not isAresId, got", lineParts.IsAresId)
	}
	if lineParts.IsBlank {
		t.Error("Expected ! isBlank, got", lineParts.IsBlank)
	}
	if lineParts.IsCommentedOut {
		t.Error("Expected ! isCommentedOut, got", lineParts.IsCommentedOut)
	}
	if !lineParts.IsRedirect {
		t.Error("Expected isRedirect, got", lineParts.IsRedirect)
	}
	if lineParts.Key != "x" {
		t.Error("Expected key = x, got", lineParts.Key)
	}

	if len(lineParts.Value) > 0 {
		t.Error("Expected len(value) < 1, got", lineParts.Value)
	}
	if lineParts.Redirect != "gr_gr_cog~actors~Priest" {
		t.Error("Expected redirect = gr_gr_cog~actors~Priest, got", lineParts.Redirect)
	}
	if len(lineParts.Comment) > 0 {
		t.Error("Expected len(comment) == 0, got", lineParts.Comment)
	}
}

// Test with key, value, and comment
func TestLineHasKVC(t *testing.T) {
	l := "x = \"y\" // I am a comment"
	lineParts, err := ParseLine(fnp, l)
	if err != nil {
		t.Error("expected no error, got", err.Error())
	}
	if !lineParts.HasComment {
		t.Error("Expected hasComment, got", lineParts.HasComment)
	}
	if !lineParts.HasValue {
		t.Error("Expected hasValue, got", lineParts.HasValue)
	}
	if lineParts.IsAresId {
		t.Error("Expected not isAresId, got", lineParts.IsAresId)
	}
	if lineParts.IsBlank {
		t.Error("Expected ! isBlank, got", lineParts.IsBlank)
	}
	if lineParts.IsCommentedOut {
		t.Error("Expected ! isCommentedOut, got", lineParts.IsCommentedOut)
	}
	if lineParts.IsRedirect {
		t.Error("Expected ! isRedirect, got", lineParts.IsRedirect)
	}
	if lineParts.Key != "x" {
		t.Error("Expected key = x, got", lineParts.Key)
	}
	if lineParts.Value != "y" {
		t.Error("Expected value = y, got", lineParts.Value)
	}
	if len(lineParts.Comment) < 1 {
		t.Error("Expected len(comment) < 0, got", lineParts.Comment)
	}
}

// Test with commented out line
func TestLineIsComment(t *testing.T) {
	l := "// x = \"y\""
	lineParts, err := ParseLine(fnp,l)
	if err != nil {
		t.Error("expected no error, got", err.Error())
	}
	if lineParts.HasComment {
		t.Error("Expected not hasComment, got", lineParts.HasComment)
	}
	if lineParts.HasValue {
		t.Error("Expected not hasValue, got", lineParts.HasValue)
	}
	if lineParts.IsAresId {
		t.Error("Expected not isAresId, got", lineParts.IsAresId)
	}
	if lineParts.IsBlank {
		t.Error("Expected ! isBlank, got", lineParts.IsBlank)
	}
	if !lineParts.IsCommentedOut {
		t.Error("Expected isCommentedOut, got", lineParts.IsCommentedOut)
	}
	if lineParts.IsRedirect {
		t.Error("Expected ! isRedirect, got", lineParts.IsRedirect)
	}
	if len(lineParts.Key) > 0 {
		t.Error("Expected len(key) < 1, got", lineParts.Key)
	}
	if len(lineParts.Value) > 0 {
		t.Error("Expected len(value) < 1, got", lineParts.Value)
	}
	if len(lineParts.Comment) > 0 {
		t.Error("Expected len(comment) < 1, got", lineParts.Comment)
	}
}

// Test with blank line
func TestLineBlank(t *testing.T) {
	l := ""
	lineParts, err := ParseLine(fnp, l)
	if err != nil {
		t.Error("expected no error, got", err.Error())
	}
	if lineParts.HasComment {
		t.Error("Expected not hasComment, got", lineParts.HasComment)
	}
	if lineParts.HasValue {
		t.Error("Expected not hasValue, got", lineParts.HasValue)
	}
	if lineParts.IsAresId {
		t.Error("Expected not isAresId, got", lineParts.IsAresId)
	}
	if !lineParts.IsBlank {
		t.Error("Expected isBlank, got", lineParts.IsBlank)
	}
	if lineParts.IsCommentedOut {
		t.Error("Expected ! isCommentedOut, got", lineParts.IsCommentedOut)
	}
	if lineParts.IsRedirect {
		t.Error("Expected ! isRedirect, got", lineParts.IsRedirect)
	}
	if len(lineParts.Key) > 0 {
		t.Error("Expected len(key) < 1, got", lineParts.Key)
	}
	if len(lineParts.Value) > 0 {
		t.Error("Expected len(value) < 1, got", lineParts.Value)
	}
	if len(lineParts.Comment) > 0 {
		t.Error("Expected len(comment) < 1, got", lineParts.Comment)
	}
}

// Test with malformed line
func TestLineMalformed(t *testing.T) {
	l := "x"
	_, err := ParseLine(fnp, l)
	if err == nil {
		t.Error("expected an error, got", "nil")
	}
}

// Test ParseAresFilename
func TestFilenameHappy(t *testing.T) {
	FilenameParts, _ := ParseAresFileName("actors_gr_GR_cog.ares")
	if len(FilenameParts.Country) < 1 {
		t.Error("Expected len(FilenameParts.country) > 0, got", FilenameParts.Country)
	}
}

// Test bad ParseAresFilename
func TestFilenameMissingAresExtension(t *testing.T) {
	_, err := ParseAresFileName("actors_gr_GR_cog")
	if strings.Index(err.Error(), "ares") < 0 {
		t.Error("Expected err msg about missing ares, got", err.Error())
	}
}

// Test bad ParseAresFilename
func TestFilenameMissingTopicOrDomain(t *testing.T) {
	_, err := ParseAresFileName("actors.ares")
	if strings.Index(err.Error(), "topic") < 0 {
		t.Error("Expected err msg about missing topic and/or domain, got", err.Error())
	}
}

// Test bad Redirect ID
func TestRedirectBad(t *testing.T) {
	_, err := ToRedirectId("actors")
	if strings.Index(err.Error(), "redirect") < 0 {
		t.Error("Expected err msg about bad redirect ID, got", err.Error())
	}
}

// TestCleanAres is test case for F.2019.005.
// The new feature is shown to be completed
// when this test case passes, i.e. the files
// in test/out have no errors in them.
func TestCleanAres(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	fmt.Println(filename)
	dir, _ := path.Split(filename)
	dir = path.Join(dir,"test")

	//in := path.Join(dir,"in","dismissals_gr_GR_cog.ares")
	in := path.Join(dir,"in","testcases_en_EN_cog.ares")
	//out := path.Join(dir,"out","dismissals_gr_GR_cog.ares")
	out := path.Join(dir,"out","testcases_en_EN_cog.ares")

	err := CleanAres(in, out)
	if err != nil {
		t.Error(err)
	}
	// len(*aresErrors) will be zero if all errors were
	// corrected by CleanAres.
	aresErrors := GetAresErrors(out)

	if len(*aresErrors) > 0 {
		t.Error(fmt.Sprintf("Expected no errors, got %d", len(*aresErrors)))
		for _, err := range *aresErrors {
			t.Error(err.Error())
		}
	}
}



