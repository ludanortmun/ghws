package internal

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/go-github/v74/github"
)

const NotFoundError = "not found"

type GitHubAPIFetcher struct {
	ghClient *github.Client
}

func NewGitHubAPIFetcher(client *github.Client) *GitHubAPIFetcher {
	return &GitHubAPIFetcher{
		ghClient: client,
	}
}

func (f *GitHubAPIFetcher) Fetch(target GitHubTarget, path string) ([]byte, error) {
	ownerAndRepo := strings.Split(target.Repository, "/")
	owner := ownerAndRepo[0]
	repo := ownerAndRepo[1]
	fullPath := fmt.Sprintf("%s/%s", target.Directory, path)

	file, _, _, err := f.ghClient.Repositories.GetContents(
		context.Background(),
		owner,
		repo,
		fullPath,
		&github.RepositoryContentGetOptions{},
	)

	if err != nil {
		// Special case for not found error
		if strings.Contains(err.Error(), NotFoundError) {
			return nil, errors.New(NotFoundError)
		}

		return nil, err
	}

	contents, err := file.GetContent()
	if err != nil {
		return nil, err
	}

	return []byte(contents), nil
}
