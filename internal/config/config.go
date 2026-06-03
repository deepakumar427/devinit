// internal/config/config.go
package config

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// ProjectConfig holds every decision the user makes.
// This struct flows through the entire scaffolding pipeline.
type ProjectConfig struct {
	ProjectName      string
	Language         string // "go", "node", "python"
	CreateGitHub     bool
	CreateDocker     bool
	CreateCI         bool
	GitHubUser       string
	GenerateAIReadme bool
	CreateK8s        bool
}

// Collect builds a ProjectConfig by reading flags first,
// then prompting interactively for anything not provided.
func Collect(cmd *cobra.Command, name string) (*ProjectConfig, error) {
	cfg := &ProjectConfig{ProjectName: name}

	// Read flags (empty string / false if not set)
	cfg.Language, _ = cmd.Flags().GetString("lang")
	cfg.CreateDocker, _ = cmd.Flags().GetBool("docker")
	cfg.CreateCI, _ = cmd.Flags().GetBool("ci")
	cfg.CreateGitHub, _ = cmd.Flags().GetBool("github")
	cfg.CreateK8s, _ = cmd.Flags().GetBool("k8s")

	// For anything missing, prompt interactively
	if cfg.Language == "" {
		prompt := &survey.Select{
			Message: "Select programming language:",
			Options: []string{"go", "node", "python"},
		}
		// survey.AskOne runs the prompt and writes the answer into &cfg.Language
		if err := survey.AskOne(prompt, &cfg.Language); err != nil {
			return nil, err
		}
	}

	// Multi-select for features the user wants
	var features []string
	featPrompt := &survey.MultiSelect{
		Message: "Select features to scaffold:",
		Options: []string{
			"GitHub repo",
			"Dockerfile",
			"GitHub Actions CI/CD",
			"Kubernetes Manifests",
		},
	}
	if err := survey.AskOne(featPrompt, &features); err != nil {
		return nil, err
	}

	// Parse multi-select results into booleans
	for _, f := range features {
		switch f {
		case "GitHub repo":
			cfg.CreateGitHub = true
		case "Dockerfile":
			cfg.CreateDocker = true
		case "GitHub Actions CI/CD":
			cfg.CreateCI = true
		case "Kubernetes Manifests":
            cfg.CreateK8s = true
        }
	}

	// If GitHub is selected, we need the username
	if cfg.CreateGitHub {
		prompt := &survey.Input{Message: "Your GitHub username:"}
		if err := survey.AskOne(prompt, &cfg.GitHubUser,
			survey.WithValidator(survey.Required)); err != nil {
			return nil, err
		}
	}

	// ---------------------------------------------------------
	// NEW: Ask the user if they want an AI-generated README
	// ---------------------------------------------------------
	aiPrompt := &survey.Confirm{
		Message: "Do you want to generate an AI-optimized README using Gemini?",
		Default: true,
	}
	if err := survey.AskOne(aiPrompt, &cfg.GenerateAIReadme); err != nil {
		return nil, err
	}

	return cfg, nil
}
