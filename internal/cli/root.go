package cli

import (
	"github.com/sajanv88/gocli/internal/adapter/db"
	"github.com/sajanv88/gocli/internal/adapter/frontend"
	"github.com/sajanv88/gocli/internal/adapter/router"
	"github.com/sajanv88/gocli/internal/app"
	"github.com/sajanv88/gocli/internal/domain"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	var spec domain.ProjectSpec

	cmd := &cobra.Command{
		Use:   "new [project-name]",
		Short: "Scaffold a new Go backend, (+ optional Db, + optional frontend)",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				spec.Name = args[0]
			}
			spec.OutputDir = "./" + spec.Name

			resolved, err := resolveSpec(spec)
			if err != nil {
				return err
			}

			uc := app.ScaffoldUseCase{
				Routers:   router.All(),
				DBs:       db.All(),
				Frontends: frontend.All(),
			}
			return uc.Execute(resolved)
		},
	}

	cmd.Flags().StringVar(&spec.ModulePath, "module", "", "Go module path, e.g. github.com/you/app")
	cmd.Flags().Var(&spec.Router, "router", "gin | chi | none")
	cmd.Flags().Var(&spec.Database, "db", "postgres | mysql | mongodb | redis | none")
	cmd.Flags().Var(&spec.Frontend, "frontend", "vite | nextjs | none")
	cmd.Flags().BoolVar(&spec.Force, "force", false, "overwrite an existing non-empty output directory")
	return cmd
}
