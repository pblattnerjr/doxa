package oslw

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"runtime"
	"testing"
)

var logger log.Logger
var logFile *os.File

func init() {
	logFile, err := os.OpenFile("test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	logger.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)
	logger.SetOutput(mw)
}

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
func TestLoadOslwResources(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	fmt.Println(filename)
	dir, _ := path.Split(filename)
	dir = path.Join(dir,"test")
	dbName := path.Join(dir,"test.db")
	err := Res2Sql(dir, dbName, &logger)
	if err != nil {
		t.Error(err.Error())
	}
}
// Returns the directory within which the caller is executing
func executionDir() string {
	_, filename, _, _ := runtime.Caller(0)
	dir, _ := path.Split(filename)
	return dir
}
