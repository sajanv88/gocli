package mongo

import (
	"embed"

	"github.com/sajanv88/gocli/internal/domain"
	"github.com/sajanv88/gocli/internal/infra"
)

//go:embed all:templates
var mongoTemplates embed.FS

type Generator struct{}

func (Generator) Name() domain.DatabaseOption { return domain.DBMongoDB }

func (Generator) Generate(spec domain.ProjectSpec) error {
	return infra.CopyTemplateFS(mongoTemplates, "templates", spec.OutputDir, spec)
}
