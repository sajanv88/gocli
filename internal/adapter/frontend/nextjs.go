package frontend

import (
	"github.com/sajanv88/gocli/internal/domain"
	"github.com/sajanv88/gocli/internal/infra"
)

type NextGenerator struct{}

func (NextGenerator) Name() domain.FrontendOption { return domain.FrontendNextJS }

func (NextGenerator) Generate(spec domain.ProjectSpec) error {
	pm, err := infra.ResolveNodePackageManager()
	if err != nil {
		return err
	}

	pmFlag := "--use-npm"
	if pm == "pnpm" {
		pmFlag = "--use-pnpm"
	}

	return infra.Run(spec.OutputDir, "npx", "create-next-app@latest",
		"frontend", "--ts", "--tailwind", "--eslint", "--app", "--src-dir",
		"--import-alias", "@/*", pmFlag, "--disable-git", "--yes",
	)
}
