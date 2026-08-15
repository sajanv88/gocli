package postgres

import (
	"embed"

	"github.com/sajanv88/gocli/internal/domain"
	"github.com/sajanv88/gocli/internal/infra"
)

//go:embed all:templates
var pgTemplates embed.FS

type Generator struct{}

func (Generator) Name() domain.DatabaseOption { return domain.DBPostgres }

func (Generator) Generate(spec domain.ProjectSpec) error {
	// writes db/postgres.go (pgx pool), .env.example with DATABASE_URL,
	// and a postgres block into docker-compose.yml
	return infra.CopyTemplateFS(pgTemplates, "templates", spec.OutputDir, spec)
}
