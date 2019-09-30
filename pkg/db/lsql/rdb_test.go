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
	repos := []string{"https://github.com/AGES-Initiatives/ages-alwb-library-gr-gr-cog.git","https://github.com/AGES-Initiatives/ages-alwb-library-en-us-dedes.git","https://github.com/AGES-Initiatives/ages-alwb-library-en-us-goa.git","https://github.com/AGES-Initiatives/ages-alwb-library-en-us-holycross.git","https://github.com/AGES-Initiatives/ages-alwb-library-en-us-oca.git","https://github.com/AGES-Initiatives/ages-alwb-library-ancillary.git","https://github.com/AGES-Initiatives/ages-alwb-library-client-enpublic.git","https://github.com/AGES-Initiatives/ages-alwb-library-en-fr-omvol.git","https://github.com/AGES-Initiatives/ages-alwb-library-en-uk-ware.git","https://github.com/AGES-Initiatives/ages-alwb-library-en-us-acook.git","https://github.com/AGES-Initiatives/ages-alwb-library-en-us-andronache.git","https://github.com/AGES-Initiatives/ages-alwb-library-en-us-barrett.git","https://github.com/AGES-Initiatives/ages-alwb-library-en-us-boyer.git","https://github.com/AGES-Initiatives/ages-alwb-library-en-us-constantinides.git","https://github.com/AGES-Initiatives/ages-alwb-library-en-us-duvall.git","https://github.com/AGES-Initiatives/ages-alwb-library-en-us-houpos.git","https://github.com/AGES-Initiatives/ages-alwb-library-en-us-repass.git","https://github.com/AGES-Initiatives/ages-alwb-library-en-us-unknown.git","https://github.com/AGES-Initiatives/ages-alwb-library-lash.git","https://github.com/AGES-Initiatives/ages-alwb-scripture.git"}
	err := Repos2Sqlite(repos, "test.db", true, &logger)
	if err != nil {
		t.Error("error:", err.Error())
	}
}
