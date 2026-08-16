package adk

import (
	"embed"

	"github.com/sajanv88/gocli/internal/domain"
	"github.com/sajanv88/gocli/internal/infra"
)

//go:embed all:templates
var adkTemplates embed.FS

type Generator struct{}

func (Generator) Name() domain.AgentOption { return domain.AgentADK }

func (Generator) Generate(spec domain.ProjectSpec) error {
	return infra.CopyTemplateFS(adkTemplates, "templates", spec.OutputDir, spec)
}
