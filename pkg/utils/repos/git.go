package repos

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// converts repository clone url into
// a path into which to clone the repository
func DirPath(parentDir, theUrl string) (string, error) {
	var err error
	Url, err := url.Parse(strings.TrimSpace(theUrl))
	result := filepath.Join(parentDir, Url.Path)
	result = result[:len(result)-4]
	return result, err
}

// Clone sequentially each repository to the parent
// directory specified in path.
// The directory will be created by the Clone function.
// Be aware that the contents of the root directory will
// be deleted if it already exists.
func CloneRepos(path string, urls []string, deleteFirst bool) error {
	_ = os.RemoveAll(path)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	for _, url := range urls {
		u, err := Clone(path, url, deleteFirst)
		fmt.Println(u)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return err
}

// Clone concurrently each repository to the parent
// directory specified in path.
// The directory will be created by the Clone function.
// Be aware that the contents of the root directory will
// be deleted if it already exists. This func runs nearly
// twice as fast as the CloneRepos func.
func CloneConcurrently(path string, urls []string) error {
	_ = os.RemoveAll(path)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	done := make(chan string, len(urls))
	errch := make(chan error, len(urls))

	for _, url := range urls {
		go func(path, url string) {
			u, err := Clone(path, url, false)
			done <- u
			errch <- err
		}(path, url)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-done)
		if err := <-errch; err != nil {
			fmt.Println(err.Error())
		}
	}
	return err
}

func Clone(path string, url string, deleteFirst bool) (string, error) {
	dirPath, err := DirPath(path, url)
	if err != nil {
		return url, err
	}
	if deleteFirst {
		err = os.RemoveAll(dirPath)
		if err != nil {
			return url, err
		}
	}
	_, err = git.PlainClone(dirPath, false, &git.CloneOptions{
		URL: url,
	})
	return dirPath, err
}

// adds *, then commits, using the msg
func Commit(path, url, msg string) error {
	dirPath, err := DirPath(path, url)
	if err != nil {
		return err
	}
	r, err := git.PlainOpen(dirPath)
	if err != nil {
		return err
	}
	w, err := r.Worktree()
	if err != nil {
		return err
	}
	_, err = w.Commit(msg, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "doxasi",
			Email: "olw@ocmc.org",
			When:  time.Now(),
		},
	})
	return err
}
func Pull(path, url string) (plumbing.Hash, error) {
	var hash plumbing.Hash
	dirPath, err := DirPath(path, url)
	if err != nil {
		return hash, err
	}
	r, err := git.PlainOpen(dirPath)
	if err != nil {
		return hash, err
	}
	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil {
		return hash, err
	}
	// Pull the latest changes from the origin remote and merge into the current branch
	err = w.Pull(&git.PullOptions{
		RemoteName: "origin",
		Force: true,
	})
	if err != nil {
		return hash, err
	}
	// Print the latest commit that was just pulled
	ref, err := r.Head()
	if err != nil {
		return hash, err
	}
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		return hash, err
	}
	return commit.Hash, err
}
func Push(path, url, username, password string) error {
	dirPath, err := DirPath(path, url)
	if err != nil {
		return err
	}
	r, err := git.PlainOpen(dirPath)
	if err != nil {
		return err
	}
	err = r.Push(&git.PushOptions{
		Auth: &http.BasicAuth{
			Username: username,
			Password: password,
		},
	})
	return err
}

// will add files of all types except for .git dir and its contents.
// TODO: add filter from gitignore file in root dir
func FilesToProcess(dirPath string) []string {
	var files []string
	filepath.Walk(dirPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.Contains(path, "/.git/") {
				return nil
			}
			if !info.IsDir() {
				files = append(files, info.Name())
			}
			return nil
		})
	return files
}
