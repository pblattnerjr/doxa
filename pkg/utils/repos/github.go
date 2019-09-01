package repos

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/liturgiko/doxa/pkg/utils/ares"
	"golang.org/x/oauth2"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"
	"log"
	"net/http"
	"strings"
)

// Get all the repositories for the supplied user name or org at github.
// Token: a github token.
// When you get back the array of repositories, you can access the
// values of repo properties using an asterisk, e.g. *repo.Name to get the repository name.
func GetAllGithubRepos(name string, token string) ([]*github.Repository, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}
	// get all pages of results
	var allRepos []*github.Repository
	for {
		repos, resp, err := client.Repositories.ListByOrg(ctx, name, opt)
		if err != nil {
			return allRepos, err
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	return allRepos, nil
}

// Clone repo into memory and iterate over its commits
func ProcessRepoCommits(name string) error {
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: name,
	})
	ref, err := r.Head()

	// ... retrieves the commit history
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})

	// ... iterates over the commits, printing them
	err = cIter.ForEach(func(c *object.Commit) error {
		fmt.Println(c)
		return nil
	})
	return err
}

// maps lines from ares files in Github repositories
// to LineParts
func Ares2LpFromGithub(repos []string,
	fileSuffix string,
	printProgress bool,
	logger *log.Logger) <-chan *ares.LineParts {
	out := make(chan *ares.LineParts)
	totalRepos := len(repos)
	reposProcessed := 0
	go func() {
		for _, repo := range repos {
			if !printProgress {
				reposProcessed++
				fmt.Print(fmt.Sprintf("Processing %d/%d: %s\n", reposProcessed, totalRepos, repo))
			}
			// test to verify we have internet connectivity and the repo exists
			_, err := http.Get(repo)
			if err != nil {
				fmt.Println(fmt.Sprintf("No Internet connection, or bad url: %s", repo))
				close(out)
			}
			// clone the repo into memory
			r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
				URL: repo,
			})

			if err != nil {
				// try again because it might have been caused by a netword delay (http request timeout)
				r, err = git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
					URL: repo,
				})
				if err != nil {
					fmt.Println(fmt.Sprintf("Error: could not clone %s", repo))
					close(out)
				}
			}

			ref, err := r.Head()
			if err != nil {
				fmt.Println(err.Error())
				close(out)
			}

			// ... retrieving the commit object
			commit, err := r.CommitObject(ref.Hash())
			if err != nil {
				fmt.Println(err.Error())
				close(out)
			}

			// ... retrieve the tree from the commit
			tree, err := commit.Tree()
			if err != nil {
				fmt.Println(err.Error())
				close(out)
			}
			// ... get the files iterator and print the file
			err = tree.Files().ForEach(func(f *object.File) error {
				if strings.HasSuffix(f.Name, fileSuffix) {
					var fileNameParts ares.FilenameParts
					fileNameParts, err := ares.ParseAresFileName(f.Name)
					if printProgress {
						fmt.Printf("\r%-80s", "")
						fmt.Print(fmt.Sprintf("\rProcessing %s_%s_%s_%s.ares", fileNameParts.Topic, fileNameParts.Language, fileNameParts.Country, fileNameParts.Realm))
					}

					var lineCnt int
					inCommentBlock := false

					if err == nil {
						l, _ := f.Lines()
						for _, line := range l {
							lineCnt = lineCnt + 1
							line = strings.TrimSpace(line)
							line = strings.TrimSpace(line)
							if strings.HasPrefix(line, "/*") {
								inCommentBlock = true
								continue
							}
							if strings.HasPrefix(line, "*/") || strings.HasSuffix(line, "*/") {
								inCommentBlock = false
								continue
							}
							if !inCommentBlock {
								lineParts, err := ares.ParseLine(fileNameParts, line)
								lineParts.LineNbr = lineCnt
								if err != nil {
									logger.Printf("%s: %s: line %d", err.Error(), f.Name, lineCnt)
									continue
								}
								if lineParts.IsAresId || lineParts.IsBlank || lineParts.IsCommentedOut {
									// ignore
								} else {
									out <- &lineParts
								}
							}
						}
					}
				}
				return err
			})
		}
		close(out)
	}()
	return out
}
