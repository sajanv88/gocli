package db

import (
	"embed"

	"github.com/sajanv88/gocli/internal/domain"
	"github.com/sajanv88/gocli/internal/infra"
)

//go:embed templates/*.tmpl
var mongoTemplates embed.FS

type MongoGenerator struct{}

func (MongoGenerator) Name() string { return "mongo" }

func (MongoGenerator) Generate(spec domain.ProjectSpec) error {
	return infra.CopyTemplateFS(mongoTemplates, "templates", spec.OutputDir, spec)
}
