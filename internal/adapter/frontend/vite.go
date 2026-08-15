package frontend

import (
	"github.com/sajanv88/gocli/internal/domain"
	"github.com/sajanv88/gocli/internal/infra"
)

type ViteGenerator struct{}

func (ViteGenerator) Name() domain.FrontendOption { return domain.FrontendVite }

func (ViteGenerator) Generate(spec domain.ProjectSpec) error {
	pm, err := infra.ResolveNodePackageManager()

	if err != nil {
		return err
	}

	if pm == "pnpm" {
		return infra.Run(spec.OutputDir, "pnpm", "create", "vite",
			"frontend",
			"--template", "react-ts",
			"--no-interactive",
		)
	}

	return infra.Run(spec.OutputDir, "npm", "create", "vite@latest",
		"frontend", "--",
		"--template", "react-ts",
		"--no-interactive",
	)
}
