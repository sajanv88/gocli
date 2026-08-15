package chi

import (
	"embed"

	"github.com/sajanv88/gocli/internal/domain"
	"github.com/sajanv88/gocli/internal/infra"
)

//go:embed templates/*.tmpl
var chiTemplates embed.FS

type Generator struct{}

func (Generator) Name() domain.RouterOption { return domain.RouterChi }

func (Generator) Generate(spec domain.ProjectSpec) error {
	return infra.CopyTemplateFS(chiTemplates, "templates", spec.OutputDir, spec)
}
