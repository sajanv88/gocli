package enio

import (
	"embed"

	"github.com/sajanv88/gocli/internal/domain"
	"github.com/sajanv88/gocli/internal/infra"
)

//go:embed all:templates
var enioTemplates embed.FS

type Generator struct{}

func (Generator) Name() domain.AgentOption { return domain.AgentEino }

func (Generator) Generate(spec domain.ProjectSpec) error {
	return infra.CopyTemplateFS(enioTemplates, "templates", spec.OutputDir, spec)
}
