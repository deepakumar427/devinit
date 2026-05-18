// internal/scaffold/docker.go
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

type DockerStep struct{}

func (d *DockerStep) Name() string { return "Dockerfile" }

func (d *DockerStep) Run(cfg *config.ProjectConfig) error {
	// Select the right template string for the language
	tmplStr, ok := dockerTemplates[cfg.Language]
	if !ok {
		tmplStr = dockerTemplates["default"]
	}

	// Parse the template
	tmpl, err := template.New("dockerfile").Parse(tmplStr)
	if err != nil {
		return fmt.Errorf("template parse error: %w", err)
	}

	// Execute template into a buffer ({{.ProjectName}} gets replaced)
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, cfg); err != nil {
		return fmt.Errorf("template execute error: %w", err)
	}

	// Write buffer contents to Dockerfile
	path := filepath.Join(cfg.ProjectName, "Dockerfile")
	return os.WriteFile(path, buf.Bytes(), 0644)
}

// dockerTemplates maps language → Dockerfile template string.
// {{.ProjectName}} and {{.Language}} are replaced at execution time.
var dockerTemplates = map[string]string{
	"go": `# Build stage
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o {{.ProjectName}} .

# Runtime stage (minimal image)
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/{{.ProjectName}} .
EXPOSE 8080
CMD ["./{{.ProjectName}}"]
`,
	"node": `FROM node:20-alpine
WORKDIR /app
COPY package*.json ./
RUN npm ci --only=production
COPY . .
EXPOSE 3000
CMD ["node", "index.js"]
`,
	"python": `FROM python:3.12-slim
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY . .
EXPOSE 8000
CMD ["python", "main.py"]
`,
	"default": `FROM ubuntu:22.04
WORKDIR /app
COPY . .
`,
}