package docker

import (
	"embed"

	"github.com/sajanv88/gocli/internal/domain"
	"github.com/sajanv88/gocli/internal/infra"
)

//go:embed all:templates
var dockerTemplates embed.FS

// Generate writes a multi-stage Dockerfile and .dockerignore for the
// scaffolded Go backend. Only called when spec.Docker is true.
func Generate(spec domain.ProjectSpec) error {
	return infra.CopyTemplateFS(dockerTemplates, "templates", spec.OutputDir, spec)
}
