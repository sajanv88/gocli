package mysql

import (
	"embed"

	"github.com/sajanv88/gocli/internal/domain"
	"github.com/sajanv88/gocli/internal/infra"
)

//go:embed all:templates
var mysqlTemplates embed.FS

type Generator struct{}

func (Generator) Name() domain.DatabaseOption { return domain.DBMySQL }

func (Generator) Generate(spec domain.ProjectSpec) error {
	// writes db/mysql.go, .env.example with DATABASE_URL,
	// and a mysql block into docker-compose.yml
	return infra.CopyTemplateFS(mysqlTemplates, "templates", spec.OutputDir, spec)
}
