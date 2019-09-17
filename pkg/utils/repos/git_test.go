package repos

import (
	"fmt"
	"os"
	"testing"
)

func TestClone(t *testing.T) {
	repos := []string{"https://github.com/liturgiko/testrepo1.git", "https://github.com/liturgiko/testrepo2.git"}
	err := CloneConcurrently("/Users/mac002/canBeRemoved/doxago/git", repos)
	if err != nil {
		t.Error("error:", err.Error())
	}
}
func TestDirPath(t *testing.T) {
	home := "/Users/mac002/canBeRemoved/doxago/git"
	url := "https://github.com/liturgiko/testrepo1.git"
	p, err := DirPath(home, url)
	if err != nil {
		t.Error("error: ", err.Error())
	}
	if p != "repos/liturgiko/testrepo1" {
		t.Error("error: invalid dir path", err.Error())
	}
}
func TestCommit(t *testing.T) {
	home := os.Getenv("dir")
	url := os.Getenv("url")
	err := Commit(home, url, "test commit")
	if err != nil {
		t.Error(err.Error())
	}
}
func TestPush(t *testing.T) {
	home := os.Getenv("dir")
	url := os.Getenv("url")
	err := Commit(home, url, "test commit")
	if err != nil {
		t.Error("commit error:", err.Error())
	}
	err = Push(home, url, os.Getenv("usr"), os.Getenv("pwd"))
	if err != nil {
		t.Error("push error", err.Error())
	}
}
func TestPull(t *testing.T) {
	home := os.Getenv("dir")
	url := os.Getenv("url")
	hash, err := Pull(home, url)
	if err != nil {
		t.Error("pull error", err.Error())
	}
	fmt.Println(hash)
}

func TestFilesToProcess(t *testing.T) {
	home := os.Getenv("dir")
	url := os.Getenv("url")
	if home == "" || url == "" {
		t.Error("home or url environment variable not set")
	} else {
		dir, _ := DirPath(home, url)
		fmt.Println(FilesToProcess(dir))
	}
}
