package db

import (
	"embed"

	"github.com/sajanv88/gocli/internal/domain"
	"github.com/sajanv88/gocli/internal/infra"
)

//go:embed templates/*.tmpl
var mysqlTemplates embed.FS

type MySQLGenerator struct{}

func (MySQLGenerator) Name() string { return "mysql" }

func (MySQLGenerator) Generate(spec domain.ProjectSpec) error {
	// writes db/mysql.go, .env.example with DATABASE_URL,
	// and a mysql block into docker-compose.yml
	return infra.CopyTemplateFS(mysqlTemplates, "templates", spec.OutputDir, spec)
}
