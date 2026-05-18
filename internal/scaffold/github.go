package scaffold

import (
	"fmt"
	"os/exec"

	"devinit/internal/config"
)

type GitHubStep struct{}

func (gh *GitHubStep) Name() string { return "GitHub repo" }

func (gh *GitHubStep) Run(cfg *config.ProjectConfig) error {
	// 1. Check if gh CLI is installed
	if _, err := exec.LookPath("gh"); err != nil {
		return fmt.Errorf("GitHub CLI (gh) not found. Please install from https://cli.github.com")
	}

	// 2. Stage all the newly generated files
	addCmd := exec.Command("git", "add", ".")
	addCmd.Dir = cfg.ProjectName // Ensure this runs inside the new project folder
	if err := addCmd.Run(); err != nil {
		return fmt.Errorf("failed to stage files: %w", err)
	}

	// 3. Create the initial commit
	commitCmd := exec.Command("git", "commit", "-m", "feat: initial commit by devinit 🚀")
	commitCmd.Dir = cfg.ProjectName // Ensure this runs inside the new project folder
	if out, err := commitCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to commit files: %s", string(out))
	}

	// 4. Create the remote repo and push
	repoName := fmt.Sprintf("%s/%s", cfg.GitHubUser, cfg.ProjectName)
	
	ghCmd := exec.Command("gh", "repo", "create", repoName,
		"--public",
		"--source", ".", // We use "." because we set the Dir below
		"--remote", "origin",
		"--push",
	)
	ghCmd.Dir = cfg.ProjectName // Ensure this runs inside the new project folder

	if out, err := ghCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("gh repo create failed: %s", string(out))
	}

	return nil
}