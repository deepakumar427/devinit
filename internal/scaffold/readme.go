package scaffold

import (
	"bytes"
	_ "embed"
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
	tmpl := template.Must(template.New("readme").Parse(readmeTmpl))

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, cfg); err != nil {
		return err
	}

	path := filepath.Join(cfg.ProjectName, "README.md")
	return os.WriteFile(path, buf.Bytes(), 0644)
}
