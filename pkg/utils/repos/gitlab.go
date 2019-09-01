package repos

import (
	"github.com/xanzy/go-gitlab"
	"log"
)

// Get a list of the Gitlab user accounts
// But, the interface requires pagination.
// So, TODO: append projects and return the complete list.
func GetUsers(url string, token string) ([]*gitlab.User, error) {
	git := gitlab.NewClient(nil, token)
	err := git.SetBaseURL(url + "/api/v4")

	if err != nil {
		err.Error()
	}
	opt := &gitlab.ListUsersOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 10,
			Page:    1,
		},
	}
	var users []*gitlab.User

	for {
		// Get the first page with projects.
		items, resp, err := git.Users.ListUsers(opt)
		if err != nil {
			log.Fatal(err)
		}

		// Append the users we've found so far.
		for _, item := range items {
			users = append(users, item)
		}

		// Exit the loop when we've seen all pages.
		if resp.CurrentPage >= resp.TotalPages {
			break
		}

		// Update the page number to get the next page.
		opt.Page = resp.NextPage
	}
	return users, err
}

// Should return an array of all projects.
// But, the interface requires pagination.
// So, TODO: append projects and return the complete list.
func GetProjects(url string, token string) ([]*gitlab.Project, error) {
	git := gitlab.NewClient(nil, token)
	err := git.SetBaseURL(url + "/api/v4")

	if err != nil {
		err.Error()
	}
	opt := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 10,
			Page:    1,
		},
	}
	var projects []*gitlab.Project

	for {
		// Get the first page with projects.
		items, resp, err := git.Projects.ListProjects(opt)
		if err != nil {
			log.Fatal(err)
		}

		// List all the projects we've found so far.
		for _, item := range items {
			projects = append(projects, item)
		}

		// Exit the loop when we've seen all pages.
		if resp.CurrentPage >= resp.TotalPages {
			break
		}

		// Update the page number to get the next page.
		opt.Page = resp.NextPage
	}
	return projects, err
}
