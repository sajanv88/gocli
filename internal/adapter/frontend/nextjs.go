package frontend

import (
	"github.com/sajanv88/gocli/internal/domain"
	"github.com/sajanv88/gocli/internal/infra"
)

type NextGenerator struct{}

func (NextGenerator) Name() string { return "nextjs" }

func (NextGenerator) Generate(spec domain.ProjectSpec) error {
	if err := infra.CheckToolAvailable("npx", "Node 20.9+"); err != nil {
		return err
	}
	return infra.Run(spec.OutputDir, "npx", "create-next-app@latest",
		"frontend", "--ts", "--tailwind", "--eslint", "--app", "--src-dir",
		"--import-alias", "@/*", "--use-npm", "--disable-git", "--yes",
	)
}
