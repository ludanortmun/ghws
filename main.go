package main

import (
	"net/http"

	"github.com/google/go-github/v74/github"
	"github.com/ludanortmun/ghws/internal"
)

func main() {

	client := github.NewClient(nil).WithAuthToken("token")

	fetcher := internal.NewGitHubAPIFetcher(client)
	server := &http.Server{
		Addr: ":8080",
	}

	ghHandler := internal.NewGitHubHandler(
		fetcher,
	).AddRootSite(internal.GitHubTarget{
		Repository: "ludanortmun/cetys-icc-hetplat",
		Directory:  "2.Presentation/www",
		Branch:     "main",
	})

	http.Handle("/{path...}", ghHandler)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
