package cmd

import (
	"log"
	"net/http"

	"github.com/google/go-github/v74/github"
	"github.com/ludanortmun/ghws/internal"
	"github.com/spf13/cobra"
)

// TODO: Add support for multiple mappings in CLI
// TODO: Add support for external config files
// TODO: Support for customizing port

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve [github url]",
	Short: "Serves the website specified by github url",
	Long: `Creates a web server to serve all the files specified in the GitHub URL.
If the URL contains a specific branch, tag, or commit, the server will only serve the files from that specified reference.
If not, the server will only serve the files from the default branch for the repository.
If the URL contains a branch (or when using the default one), the server will always respond with the latest version of the files at that branch.
If the URL contains a tag or commit hash, the files served will be pinned to that revision.
`,
	Run: func(cmd *cobra.Command, args []string) {
		startServer(args[0])
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func startServer(url string) {

	client := github.NewClient(nil)

	fetcher := internal.NewGitHubAPIFetcher(client)
	server := &http.Server{
		Addr: ":8080",
	}

	target, err := internal.InferTargetFromUrl(url)
	if err != nil {
		log.Fatal(err)
	}

	ghHandler := internal.NewGitHubHandler(
		fetcher,
	).AddRootSite(target)

	http.Handle("/{path...}", ghHandler)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
