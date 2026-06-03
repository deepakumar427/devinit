// internal/scaffold/scaffold.go
package scaffold

import (
	"fmt"
	"os"

	"github.com/fatih/color"

	// Corrected to match your local go.mod initialization
	"devinit/internal/config"
)

// Step is an interface every scaffold action must implement.
// This is the key Go design pattern: program to interfaces.
type Step interface {
	Name() string
	Run(cfg *config.ProjectConfig) error
}

func Run(cfg *config.ProjectConfig) error {
	// Create the project directory first
	if err := os.MkdirAll(cfg.ProjectName, 0755); err != nil {
		return fmt.Errorf("cannot create directory: %w", err)
	}

	// 1. Git Init
	steps := []Step{&GitStep{}}
	
	// 2. Generate all files FIRST
	steps = append(steps, &ReadmeStep{})
	if cfg.CreateDocker {
		steps = append(steps, &DockerStep{})
	}
	if cfg.CreateCI {
		steps = append(steps, &CICDStep{})
	}

	// 3. Commit and push to GitHub LAST
	if cfg.CreateGitHub {
		steps = append(steps, &GitHubStep{})
	}

	// Execute each step, stopping on first error
	for _, step := range steps {
		color.Cyan("  [WAIT] %s...", step.Name())
		if err := step.Run(cfg); err != nil {
			color.Red("  [FAIL] %s failed: %v", step.Name(), err)
			return err
		}
		color.Green("  [OK]   %s", step.Name())
	}

	return nil
}