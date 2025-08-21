package internal

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/go-github/v74/github"
)

const NotFoundError = "not found"

type ApiFetcher struct {
	client *github.Client
}

func NewGitHubAPIFetcher(client *github.Client) *ApiFetcher {
	return &ApiFetcher{
		client: client,
	}
}

func (f *ApiFetcher) Fetch(target GitHubTarget, path string) ([]byte, error) {
	fullPath := fmt.Sprintf("%s/%s", target.directory, strings.TrimPrefix(path, "/"))

	opts := &github.RepositoryContentGetOptions{
		Ref: target.ref,
	}

	file, _, _, err := f.client.Repositories.GetContents(
		context.Background(),
		target.owner,
		target.repository,
		fullPath,
		opts,
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
