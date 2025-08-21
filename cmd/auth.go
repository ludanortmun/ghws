package cmd

import (
	"syscall"

	"github.com/ludanortmun/ghws/internal"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// authCmd represents the config command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Allows for configuring authentication for the CLI application",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

var setTokenCmd = &cobra.Command{
	Use:   "set-token",
	Short: "Sets the GitHub Personal Access Token (PAT) for authenticated API access",
	Long: `Sets the GitHub Personal Access Token (PAT) for authenticated API access.
This command will store the token securely using the system's keyring.`,
	Run: func(cmd *cobra.Command, args []string) {
		setGithubToken()
	},
}

var clearTokenCmd = &cobra.Command{
	Use:   "clear-token",
	Short: "Clears the stored GitHub Personal Access Token (PAT)",
	Long: `Clears the stored GitHub Personal Access Token (PAT).
This command will remove the token from the system's keyring.`,
	Run: func(cmd *cobra.Command, args []string) {
		clearGithubToken()
	},
}

func init() {
	authCmd.AddCommand(setTokenCmd)
	authCmd.AddCommand(clearTokenCmd)
	rootCmd.AddCommand(authCmd)
}

func setGithubToken() {
	println("Enter your GitHub Personal Access Token (PAT):")
	tokenBytes, err := term.ReadPassword(int(syscall.Stdin))

	if err != nil {
		println("Error reading token:", err.Error())
		return
	}

	token := string(tokenBytes)

	err = internal.SaveAuthToken(token)
	if err != nil {
		println("Error saving token:", err.Error())
		return
	}

	println("GitHub PAT saved successfully.")
}

func clearGithubToken() {
	_ = internal.DeleteAuthToken()
	println("GitHub PAT cleared successfully.")
}
