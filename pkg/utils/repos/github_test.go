package repos

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestReposToFileLines(t *testing.T) {
	var logger log.Logger
	var logFile *os.File
	logFile, err := os.OpenFile("github_test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	logger.SetOutput(logFile)
	logger.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)

	repos := []string{"https://github.com/liturgiko/testrepo1.git", "https://github.com/liturgiko/testrepo2.git"}
	for line := range Ares2LpFromGithub(repos, "ares", true, &logger) {
		fmt.Println(line)
	}
}

// Test getting all Github repos
// To run this test, you have to supply a valid github token
// by setting an environmental variable named githubtoken.
func TestGithubRepos(t *testing.T) {
	name := "AGES-Initiatives"
	repos, err := GetAllGithubRepos(name, os.Getenv("githubToken"))
	// Print out info about each repo:
	for _, repo := range repos {
		fmt.Printf("- %s\n", *repo.CloneURL)
	}

	if err != nil {
		t.Error("error:", err.Error())
	}
	if len(repos) < 1 {
		t.Error("error: no repos for ", name)
	}
}

// Test in-memory clone of repo from Github.
func TestGitClone(t *testing.T) {
	repo := "https://github.com/mcolburn/aresTest.git"
	err := ProcessRepoCommits(repo)
	if err != nil {
		t.Error("error:", err.Error())
	}
}
