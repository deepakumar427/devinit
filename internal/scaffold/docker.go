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

//go:embed templates/dockerfile.tmpl
var dockerfileTmpl string

type DockerStep struct{}

func (d *DockerStep) Name() string { return "Dockerfile" }

func (d *DockerStep) Run(cfg *config.ProjectConfig) error {
	tmpl, err := template.New("dockerfile").Parse(dockerfileTmpl)
	if err != nil {
		return fmt.Errorf("template parse error: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, cfg); err != nil {
		return fmt.Errorf("template execute error: %w", err)
	}

	path := filepath.Join(cfg.ProjectName, "Dockerfile")
	return os.WriteFile(path, buf.Bytes(), 0644)
}
