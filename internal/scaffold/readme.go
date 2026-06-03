package scaffold

import (
	"bytes"
	_ "embed" // Fixed: Added the blank identifier for the embed package
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"devinit/internal/config"
)

//go:embed templates/readme.tmpl
var readmeTmpl string

type ReadmeStep struct{}

func (r *ReadmeStep) Name() string { return "README.md" }

func (r *ReadmeStep) Run(cfg *config.ProjectConfig) error {
	path := filepath.Join(cfg.ProjectName, "README.md")

	// 1. Attempt AI Generation First (if requested)
	if cfg.GenerateAIReadme {
		// Updated to match your new professional logging format
		fmt.Println("  [WAIT] Generating intelligent README via Gemini...")
		aiContent, err := GenerateAIReadme(cfg)

		if err == nil {
			// Success! Write the AI content and exit the step
			return os.WriteFile(path, []byte(aiContent), 0644)
		}

		// If AI fails (e.g., missing API key), don't crash.
		// Print a warning and fall through to the standard template.
		fmt.Printf("  [FAIL] AI generation failed: %v\n  [WAIT] Falling back to standard template...\n", err)
	}

	// 2. Standard Template Fallback
	// Fixed: Actually parse the embedded template string!
	tmpl, err := template.New("readme").Parse(readmeTmpl)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, cfg); err != nil {
		return err
	}

	return os.WriteFile(path, buf.Bytes(), 0644)
}
