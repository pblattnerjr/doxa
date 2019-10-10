package oslw

import (
	"fmt"
	"testing"
)

func TestParseOslwResource(t *testing.T) {
	var id = "\\itId{en}{uk}{lash}{actors}{Bishop}{"
	var val = "Bishop"
	lineParts, err := ParseOslwResource(id, val)
	if err != nil {
		t.Error(err)
	}
	if lineParts.Language != "en" {
		t.Error(fmt.Sprintf("expected language = 'en', but got %s", lineParts.Language))
	}
	if lineParts.Country != "uk" {
		t.Error(fmt.Sprintf("expected country = 'uk', but got %s", lineParts.Country))
	}
	if lineParts.Realm != "lash" {
		t.Error(fmt.Sprintf("expected realm = 'lash', but got %s", lineParts.Realm))
	}
	if lineParts.Topic != "actors" {
		t.Error(fmt.Sprintf("expected topic = 'actors', but got %s", lineParts.Topic))
	}
	if lineParts.Value != "Bishop" {
		t.Error(fmt.Sprintf("expected value = 'Bishop', but got %s", lineParts.Value))
	}
}
// TODO: how to use a path relative to the test case
func TestLoadOslwResources(t *testing.T) {
	dir := "/Users/mac002/git/liturgika/doxa/pkg/utils/oslw/test"
	dbName := "/Users/mac002/git/liturgika/doxa/pkg/utils/oslw/test/test.db"
	err := LoadOslwResources(dir, dbName)
	if err != nil {
		t.Error(err.Error())
	}
}