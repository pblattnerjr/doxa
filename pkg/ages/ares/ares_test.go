package ares

import (
	"flag"
	"fmt"
	"github.com/liturgiko/doxa/pkg/utils/ltfile"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var fnp FilenameParts
var suppressComments *bool
var inArg *string
var outArg *string

func init() {
	fnp, _ = ParseAresFileName("actors_gr_gr_cog.ares")
	suppressComments = flag.Bool("suppressComments", false, "suppress explanatory comments")
	inArg = flag.String("in", "ABC", "input file or folder")
	outArg = flag.String("out", "XYZ", "output file or folder")
}

// Test with key and value
func TestLineHasKV(t *testing.T) {
	l := "x = \"y\""
	lineParts, err := ParseLine(fnp, l)
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
	if lineParts.Redirect != "gr_gr_cog/actors/Priest" {
		t.Error("Expected redirect = gr_gr_cog/actors/Priest, got", lineParts.Redirect)
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

var allErrors []error
var aresErrors *[]error

func TestCleanAres(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	fmt.Println(filename)
	dir, _ := filepath.Split(filename)
	dir = filepath.Join(dir, "test")

	flag.Parse()

	var in, out string

	if strings.ContainsAny(*inArg, ":/\\") { // full path
		in = *inArg
	} else {
		in = filepath.Join(dir, "in", *inArg)
	}
	InBase = in

	if strings.ContainsAny(*outArg, ":/\\") { // full path
		out = *outArg
	} else {
		out = filepath.Join(dir, "out", *outArg)
	}
	OutBase = out

	if strings.ContainsAny(*outArg, ":/\\") { // full path
		out = *outArg
	} else {
		out = filepath.Join(dir, "out", *outArg)
	}

	files, err := ltfile.FileMatcher(InBase, "ares", nil)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(files) == 0 {
		t.Fatalf("no ares files found in %s", in)
	}

	var walkErr error
	for _, file := range files {
		fileOut := filepath.Join(OutBase, filepath.Base(file))
		walkErr = CleanAres(file, fileOut, *suppressComments) // processing a file, nor a directory
		if walkErr != nil {
			allErrors = append(allErrors, walkErr)
		} else {
			aresErrors = GetAresErrors(fileOut)
			for _, err := range *aresErrors {
				allErrors = append(allErrors, err)
			}
		}
	}
	var expectedBadReferences = map[string]string{
		"gr_gr_cog":"dis101",
		"en_en_cog":"testcase9",
	}

	for libName, libRefs := range AllReferences {
		for keyName, references := range libRefs {
			if _, ok := AllDefinitions[libName][keyName]; !ok {
				if value, ok := expectedBadReferences[libName]; !ok || value != keyName {
					t.Error(fmt.Sprintf("In library %v, definition NOT found for key: %q", libName, keyName))
					for _, ref := range references {
						t.Error(fmt.Sprintf("    key referenced on line %d of %v", ref.line, ref.file))
					}
				}
			}
		}
	}

	// len(allErrors) will be zero if all errors were corrected by CleanAres.
	if len(allErrors) > 0 {
		t.Error(fmt.Sprintf("Expected no errors, got %d", len(allErrors)))
		for _, err := range allErrors {
			t.Error(err.Error())
		}
	}
}

