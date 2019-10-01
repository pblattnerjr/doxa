package lsql

import (
	"io"
	"log"
	"os"
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


// Test process Git repos
func TestRepos2Sqlite(t *testing.T) {
	repos := []string{"https://github.com/liturgiko/testrepo1.git","https://github.com/liturgiko/testrepo2.git"}
	err := Repos2Sqlite(repos, "test.db", true, &logger)
	if err != nil {
		t.Error("error:", err.Error())
	}
}
