package redis

import (
	"embed"

	"github.com/sajanv88/gocli/internal/domain"
	"github.com/sajanv88/gocli/internal/infra"
)

//go:embed all:templates
var redisTemplates embed.FS

type Generator struct{}

func (Generator) Name() domain.DatabaseOption { return domain.DBRedis }

func (Generator) Generate(spec domain.ProjectSpec) error {
	return infra.CopyTemplateFS(redisTemplates, "templates", spec.OutputDir, spec)
}
