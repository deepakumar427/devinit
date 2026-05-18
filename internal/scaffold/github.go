// internal/scaffold/github.go
package scaffold

import (
	"fmt"
	"os/exec"

	// Corrected to match your local go.mod initialization
	"devinit/internal/config"
)

type GitHubStep struct{}

func (gh *GitHubStep) Name() string { return "GitHub repo" }

func (gh *GitHubStep) Run(cfg *config.ProjectConfig) error {
	// Requires the `gh` CLI to be installed and authenticated.
	// Check if gh is available first.
	if _, err := exec.LookPath("gh"); err != nil {
		return fmt.Errorf("GitHub CLI (gh) not found. Install from https://cli.github.com")
	}

	repoName := fmt.Sprintf("%s/%s", cfg.GitHubUser, cfg.ProjectName)

	// gh repo create creates a remote repo and links it
	cmd := exec.Command("gh", "repo", "create", repoName,
		"--public",
		"--source", cfg.ProjectName,
		"--remote", "origin",
		"--push",
	)

	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("gh repo create failed: %s", string(out))
	}

	return nil
}