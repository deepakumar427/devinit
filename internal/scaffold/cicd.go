// internal/scaffold/cicd.go
package scaffold

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	// Corrected to match your local go.mod initialization
	"devinit/internal/config"
)

type CICDStep struct{}

func (c *CICDStep) Name() string { return "GitHub Actions CI/CD" }

func (c *CICDStep) Run(cfg *config.ProjectConfig) error {
	// GitHub Actions files live in .github/workflows/
	workflowDir := filepath.Join(cfg.ProjectName, ".github", "workflows")
	if err := os.MkdirAll(workflowDir, 0755); err != nil {
		return fmt.Errorf("cannot create workflow dir: %w", err)
	}

	tmplStr := ciTemplates[cfg.Language]
	if tmplStr == "" {
		tmplStr = ciTemplates["default"]
	}

	tmpl, err := template.New("ci").Parse(tmplStr)
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

var ciTemplates = map[string]string{
	"go": `name: Go CI

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.22'
      - name: Build
        run: go build ./...
      - name: Test
        run: go test ./... -v
`,
	"node": `name: Node CI

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: '20'
          cache: 'npm'
      - run: npm ci
      - run: npm test
`,
	"python": `name: Python CI

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-python@v5
        with:
          python-version: '3.12'
      - run: pip install -r requirements.txt
      - run: pytest
`,
	"default": `name: CI
on: [push]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
`,
}