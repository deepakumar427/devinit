package scaffold

import (
	"context"
	"fmt"
	"os"

	"devinit/internal/config"
	"google.golang.org/genai"
)

// GenerateAIReadme queries Gemini to write a custom README based on project context.
func GenerateAIReadme(cfg *config.ProjectConfig) (string, error) {
	// 1. Handle API Key securely via environment variable
	apiKey := os.Getenv("DEVINIT_GEMINI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("DEVINIT_GEMINI_API_KEY environment variable is not set")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return "", fmt.Errorf("failed to create AI client: %w", err)
	}

	// 2. Construct a highly specific prompt using your ProjectConfig
	prompt := fmt.Sprintf(`
You are an expert developer. Write a highly professional, concise README.md for a new project.
Project Name: %s
Primary Language: %s
Uses Docker: %t
Uses GitHub Actions CI/CD: %t

Requirements:
- Include an architecture overview, prerequisites, and a Quickstart section.
- Provide the exact terminal commands to run the project locally.
- Output ONLY the raw Markdown text. Do not wrap the entire response in markdown code blocks.`,
		cfg.ProjectName, cfg.Language, cfg.CreateDocker, cfg.CreateCI)

	// 3. Call the model
	parts := []*genai.Part{
		{Text: prompt},
	}
	
	// gemini-2.5-flash is ideal here: highly capable and very fast for CLI tools
	result, err := client.Models.GenerateContent(ctx, "gemini-2.5-flash", []*genai.Content{{Parts: parts}}, nil)
	if err != nil {
		return "", fmt.Errorf("AI generation failed: %w", err)
	}

	// 4. Extract the response safely
	if len(result.Candidates) > 0 {
		candidate := result.Candidates[0]

		// Check if the response was blocked or dropped (e.g., by safety filters)
		if candidate.Content == nil {
			return "", fmt.Errorf("AI response was blocked or empty (Reason: %v)", candidate.FinishReason)
		}

		if len(candidate.Content.Parts) > 0 && candidate.Content.Parts[0] != nil {
			return candidate.Content.Parts[0].Text, nil
		}
	}

	return "", fmt.Errorf("received empty response from AI")
}