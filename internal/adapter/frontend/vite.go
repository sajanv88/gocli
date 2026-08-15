package frontend

import (
	"github.com/sajanv88/gocli/internal/domain"
	"github.com/sajanv88/gocli/internal/infra"
)

type ViteGenerator struct{}

func (ViteGenerator) Name() string { return "vite" }

func (ViteGenerator) Generate(spec domain.ProjectSpec) error {
	if err := infra.CheckToolAvailable("npm", "Node 20.19+/22.12+"); err != nil {
		return err
	}

	return infra.Run(spec.OutputDir, "npm", "create", "vite@latest",
		"frontend", "--",
		"--template", "react-ts",
		"--no-interactive",
	)
}
