package db

import (
	"embed"

	"github.com/sajanv88/gocli/internal/domain"
	"github.com/sajanv88/gocli/internal/infra"
)

//go:embed templates/*.tmpl
var pgTemplates embed.FS

type PostgresGenerator struct{}

func (PostgresGenerator) Name() string { return "postgres" }

func (PostgresGenerator) Generate(spec domain.ProjectSpec) error {
	// writes db/postgres.go (pgx pool), .env.example with DATABASE_URL,
	// and a postgres block into docker-compose.yml
	return infra.CopyTemplateFS(pgTemplates, "templates", spec.OutputDir, spec)
}
