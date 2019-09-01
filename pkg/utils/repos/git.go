package repos

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"net/url"
	"os"
	"path/filepath"
)
// converts repository clone url into
// a path into which to clone the repository
func DirPath(parentDir, theUrl string) (string, error) {
	var err error
	Url, err := url.Parse(theUrl)
	result := filepath.Join(parentDir, Url.Path)
	result = result[:len(result)-4]
	return result, err
}
// Clone sequentially each repository to the parent
// directory specified in path.
// The directory will be created by the Clone function.
// Be aware that the contents of the root directory will
// be deleted if it already exists.
func CloneRepos(path string, urls []string) error {
	_ = os.RemoveAll(path)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	for _, url := range urls {
		u, err := Clone(path, url)
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
			u, err := Clone(path, url)
			//if err != nil {
			//	errch <- err
			//	done <- ""
			//}
			done <- u
			errch <- err
		}(path, url)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<- done)
		if err := <-errch; err != nil {
			fmt.Println(err.Error())
		}
	}
	return err
}

func Clone(path string, url string) (string, error) {
	dirPath, err := DirPath(path, url)
	if err != nil {
		return url, err
	}
	_, err = git.PlainClone(dirPath, false, &git.CloneOptions{
		URL:  url,
		//		  Progress: os.Stdout,
	})
	return dirPath, err
}
