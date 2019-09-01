package repos

import (
	"os"
	"testing"
)

// Test getting all Gitlab users
func TestGitlabUsersList(t *testing.T) {
	url := "https://gitlab.liml.org"
	items, err := GetUsers(url, os.Getenv("gitlabToken"))

	if err != nil {
		t.Error("error:", err.Error())
	}
	if len(items) < 1 {
		t.Error("error: no items found")
	}
}

// Test getting all Gitlab projects
func TestGitlabProjectsList(t *testing.T) {
	url := "https://gitlab.liml.org"
	items, err := GetProjects(url, os.Getenv("gitlabToken"))

	if err != nil {
		t.Error("error:", err.Error())
	}
	if len(items) < 1 {
		t.Error("error: no items found")
	}
}
