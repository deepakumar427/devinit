// cmd/init.go
package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	// Corrected to match your local go.mod initialization
	"devinit/internal/config"
	"devinit/internal/scaffold"
)

// initCmd represents: devinit init <project-name>
var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new developer project",
	Args:  cobra.ExactArgs(1), // requires exactly one argument
	RunE:  runInit,            // RunE returns an error (vs Run which panics)
}

func init() {
	// attach initCmd to rootCmd so `devinit init` works
	rootCmd.AddCommand(initCmd)

	// Optional flags: devinit init my-app --lang go --docker --ci
	initCmd.Flags().String("lang", "", "Programming language (go, node, python)")
	initCmd.Flags().Bool("docker", false, "Generate a Dockerfile")
	initCmd.Flags().Bool("ci", false, "Generate GitHub Actions CI/CD pipeline")
	initCmd.Flags().Bool("github", false, "Create a GitHub repository")
}

func runInit(cmd *cobra.Command, args []string) error {
	projectName := args[0]

	color.Cyan("\n[devinit] Initializing project: %s\n\n", projectName)

	// Step 1: collect configuration (via prompts or flags)
	cfg, err := config.Collect(cmd, projectName)
	if err != nil {
		return fmt.Errorf("configuration failed: %w", err)
	}
	//fmt.Printf("DEBUG: Language is: '%s'\n", cfg.Language)

	// Step 2: run all scaffolding steps
	if err := scaffold.Run(cfg); err != nil {
		return fmt.Errorf("scaffolding failed: %w", err)
	}

	color.Green("\n✓ Project '%s' is ready! cd ./%s\n", projectName, projectName)
	return nil
}
