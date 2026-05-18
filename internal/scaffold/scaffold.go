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

// Run executes all applicable scaffold steps in order.
func Run(cfg *config.ProjectConfig) error {
	// Create the project directory first
	if err := os.MkdirAll(cfg.ProjectName, 0755); err != nil {
		return fmt.Errorf("cannot create directory: %w", err)
	}

	// Build the list of steps to run based on config
	steps := []Step{
		&GitStep{},    // always runs
		&ReadmeStep{}, // always runs
	}

	if cfg.CreateGitHub {
		steps = append(steps, &GitHubStep{})
	}
	if cfg.CreateDocker {
		steps = append(steps, &DockerStep{})
	}
	if cfg.CreateCI {
		steps = append(steps, &CICDStep{})
	}

	// Execute each step, stopping on first error
	for _, step := range steps {
		color.Blue("  → %s...", step.Name())
		if err := step.Run(cfg); err != nil {
			color.Red("  ✗ %s failed: %v", step.Name(), err)
			return err
		}
		color.Green("  ✓ %s", step.Name())
	}

	return nil
}