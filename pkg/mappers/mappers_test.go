package mappers

import (
	"fmt"
	"github.com/liturgiko/doxa/pkg/utils/repos"
	"log"
	"os"
	"testing"
)

// Test that combines two channels as a pipeline:
// First we use R2Lp to map the lines in files in Github repositories to LineParts,
// then we map lineparts to Ltext structs.
func TestLp2Lt(t *testing.T) {
	var logger log.Logger
	var logFile *os.File
	logFile, err := os.OpenFile("github_test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	logger.SetOutput(logFile)
	logger.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)
	theRepos := []string{"https://github.com/liturgiko/testrepo1.git", "https://github.com/liturgiko/testrepo2.git"}
	for line := range Lp2Lt(repos.Ares2LpFromGithub(theRepos, "", true, &logger)) {
		fmt.Println(line)
	}

}
