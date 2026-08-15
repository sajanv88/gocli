package app

import (
	"github.com/sajanv88/gocli/internal/adapter/base"
	"github.com/sajanv88/gocli/internal/domain"
	"github.com/sajanv88/gocli/internal/infra"
)

type ScaffoldUseCase struct {
	Routers   map[domain.RouterOption]domain.RouterGenerator
	DBs       map[domain.DatabaseOption]domain.DBGenerator
	Frontends map[domain.FrontendOption]domain.FrontendGenerator
}

func (u ScaffoldUseCase) Execute(spec domain.ProjectSpec) error {
	if err := spec.Validate(); err != nil {
		return err
	}

	if err := infra.EnsureCleanOutputDir(spec.OutputDir, spec.Force); err != nil {
		return err
	}

	if err := infra.InitGoModule(spec.OutputDir, spec.ModulePath); err != nil {
		return err
	}

	if err := base.Generate(spec); err != nil {
		return err
	}

	if spec.Router != "none" {
		if gen, ok := u.Routers[spec.Router]; ok {
			if err := gen.Generate(spec); err != nil {
				return err
			}
		}
	}

	if spec.Database != "none" {
		if gen, ok := u.DBs[spec.Database]; ok {
			if err := gen.Generate(spec); err != nil {
				return err
			}
		}
	}

	if spec.Frontend != "none" {
		if gen, ok := u.Frontends[spec.Frontend]; ok {
			if err := gen.Generate(spec); err != nil {
				return err
			}
		}
	}

	if err := infra.Run(spec.OutputDir, "go", "mod", "tidy"); err != nil {
		return err
	}
	return infra.Run(spec.OutputDir, "git", "init")

}
