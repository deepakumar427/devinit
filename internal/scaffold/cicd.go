package scaffold

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"devinit/internal/config"
)

//go:embed templates/cicd.tmpl
var cicdTmpl string

type CICDStep struct{}

func (c *CICDStep) Name() string { return "GitHub Actions CI/CD" }

func (c *CICDStep) Run(cfg *config.ProjectConfig) error {
	workflowDir := filepath.Join(cfg.ProjectName, ".github", "workflows")
	if err := os.MkdirAll(workflowDir, 0755); err != nil {
		return fmt.Errorf("cannot create workflow dir: %w", err)
	}

	tmpl, err := template.New("ci").Parse(cicdTmpl)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, cfg); err != nil {
		return err
	}

	path := filepath.Join(workflowDir, "ci.yml")
	return os.WriteFile(path, buf.Bytes(), 0644)
}
