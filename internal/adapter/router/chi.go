package router

import (
	"embed"

	"github.com/sajanv88/gocli/internal/domain"
	"github.com/sajanv88/gocli/internal/infra"
)

//go:embed templates/*.tmpl
var chiTemplates embed.FS

type ChiGenerator struct{}

func (ChiGenerator) Name() string { return "chi" }

func (ChiGenerator) Generate(spec domain.ProjectSpec) error {
	return infra.CopyTemplateFS(chiTemplates, "templates", spec.OutputDir, spec)
}
