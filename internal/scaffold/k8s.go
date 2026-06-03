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

//go:embed templates/k8s.tmpl
var k8sTmpl string

type K8sStep struct{}

func (k *K8sStep) Name() string { return "Kubernetes Manifests" }

func (k *K8sStep) Run(cfg *config.ProjectConfig) error {
	k8sDir := filepath.Join(cfg.ProjectName, "k8s")
	if err := os.MkdirAll(k8sDir, 0755); err != nil {
		return fmt.Errorf("cannot create k8s directory: %w", err)
	}

	tmpl, err := template.New("k8s").Parse(k8sTmpl)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, cfg); err != nil {
		return err
	}

	path := filepath.Join(k8sDir, "deployment.yaml")
	return os.WriteFile(path, buf.Bytes(), 0644)
}