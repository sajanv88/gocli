package base

import (
	"embed"

	"github.com/sajanv88/gocli/internal/domain"
	"github.com/sajanv88/gocli/internal/infra"
)

//go:embed all:templates
var baseTemplates embed.FS

// Generate writes the project-wide files that every scaffold gets
// regardless of router/db/frontend choice: .gitignore and README.md.
func Generate(spec domain.ProjectSpec) error {
	return infra.CopyTemplateFS(baseTemplates, "templates", spec.OutputDir, spec)
}
