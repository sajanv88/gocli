package db

import (
	"embed"

	"github.com/sajanv88/gocli/internal/domain"
	"github.com/sajanv88/gocli/internal/infra"
)

//go:embed templates/*.tmpl
var redisTemplates embed.FS

type RedisGenerator struct{}

func (RedisGenerator) Name() string { return "redis" }

func (RedisGenerator) Generate(spec domain.ProjectSpec) error {
	return infra.CopyTemplateFS(redisTemplates, "templates", spec.OutputDir, spec)
}
