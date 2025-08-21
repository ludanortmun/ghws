package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ghws",
	Short: "Serve websites from GitHub repos.",
	Long: `GHWS is a CLI library that can be used to serve static websites from GitHub repositories.

It supports features such as:
- Selecting a specific directory as the website content root
- Supports all public repos
- Serve websites from private repositories (requires a PAT)
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
