package internal

import "fmt"

type GitHubTarget struct {
	Repository string
	Directory  string
	Branch     string
	CommitHash string
}

func (t *GitHubTarget) ToPathBase() string {
	if t.CommitHash != "" {
		return fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s", t.Repository, t.CommitHash, t.Directory)
	}
	return fmt.Sprintf("https://raw.githubusercontent.com/%s/refs/heads/%s/%s", t.Repository, t.Branch, t.Directory)
}
