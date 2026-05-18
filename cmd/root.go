// cmd/root.go
package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

// rootCmd is the base command when called without any subcommands.
// Running just `devinit` prints help.
var rootCmd = &cobra.Command{
    Use:   "devinit",
    Short: "Scaffold a production-ready developer environment in one command",
    Long: `devinit bootstraps your project with:
  - Git repository initialization
  - GitHub remote repo creation
  - Dockerfile (language-specific)
  - GitHub Actions CI/CD pipeline
  - Auto-generated README`,
}

// Execute is called by main.go. Any subcommand errors propagate here.
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}