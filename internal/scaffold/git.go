// internal/scaffold/git.go
package scaffold

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	// Corrected to match your local go.mod initialization
	"devinit/internal/config"
)

type GitStep struct{}

func (g *GitStep) Name() string { return "Git init" }

func (g *GitStep) Run(cfg *config.ProjectConfig) error {
	dir := cfg.ProjectName

	// Run `git init` inside the project directory
	// exec.Command builds the command; .CombinedOutput() runs it and captures output
	cmd := exec.Command("git", "init", dir)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git init failed: %s", out)
	}

	// Write a .gitignore based on language
	gitignore := gitignoreFor(cfg.Language)
	path := filepath.Join(dir, ".gitignore")
	if err := os.WriteFile(path, []byte(gitignore), 0644); err != nil {
		return fmt.Errorf("cannot write .gitignore: %w", err)
	}

	return nil
}

// gitignoreFor returns a sensible .gitignore for each language.
func gitignoreFor(lang string) string {
	switch lang {
	case "go":
		return "# Go\n*.exe\n*.test\n*.out\nvendor/\n"
	case "node":
		return "# Node\nnode_modules/\ndist/\n.env\n.DS_Store\n"
	case "python":
		return "# Python\n__pycache__/\n*.pyc\n.venv/\n.env\n*.egg-info/\n"
	default:
		return ".DS_Store\n.env\n"
	}
}