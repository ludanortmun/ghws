package internal

import (
	"context"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/google/go-github/v74/github"
)

const NotFoundError = "not found"

// Error message from the GitHub API when a file is not found inside a directory that does exist
var notFoundRegex = regexp.MustCompile(`no file named .+ found in .+`)

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

	c, res, err := f.client.Repositories.DownloadContents(
		context.Background(),
		target.owner,
		target.repository,
		fullPath,
		opts)

	if err != nil {
		// Special case for not found error
		if isNotFoundError(err) {
			return nil, errors.New(NotFoundError)
		}

		return nil, err
	}

	if res.StatusCode > 299 {
		if res.StatusCode == 404 {
			return nil, errors.New(NotFoundError)
		}
		return nil, fmt.Errorf("failed to fetch content: %s", res.Status)
	}

	content, err := io.ReadAll(c)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func isNotFoundError(err error) bool {
	// The regex catches errors returned by GitHub when the directory exists but the file does not.
	// The strings.Contains part catches errors when the directory itself does not exist.
	return notFoundRegex.MatchString(err.Error()) || strings.Contains(err.Error(), "Not Found")
}
