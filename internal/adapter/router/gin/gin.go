package gin

import (
	"embed"

	"github.com/sajanv88/gocli/internal/domain"
	"github.com/sajanv88/gocli/internal/infra"
)

//go:embed templates/*.tmpl
var ginTemplates embed.FS

type Generator struct{}

func (Generator) Name() domain.RouterOption { return domain.RouterGin }

func (Generator) Generate(spec domain.ProjectSpec) error {
	return infra.CopyTemplateFS(ginTemplates, "templates", spec.OutputDir, spec)
}
