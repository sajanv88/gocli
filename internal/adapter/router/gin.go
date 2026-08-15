package router

import (
	"embed"

	"github.com/sajanv88/gocli/internal/domain"
	"github.com/sajanv88/gocli/internal/infra"
)

//go:embed templates/*.tmpl
var ginTemplates embed.FS

type GinGenerator struct{}

func (GinGenerator) Name() string { return "gin" }

func (GinGenerator) Generate(spec domain.ProjectSpec) error {
	return infra.CopyTemplateFS(ginTemplates, "templates", spec.OutputDir, spec)
}
