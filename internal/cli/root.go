package cli

import (
	"fmt"
	"path/filepath"

	"github.com/sajanv88/gocli/internal/adapter/db"
	"github.com/sajanv88/gocli/internal/adapter/frontend"
	"github.com/sajanv88/gocli/internal/adapter/router"
	"github.com/sajanv88/gocli/internal/app"
	"github.com/sajanv88/gocli/internal/domain"
	"github.com/sajanv88/gocli/internal/infra"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	var spec domain.ProjectSpec
	var installDeps bool

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
			if err := uc.Execute(resolved); err != nil {
				return err
			}
			if resolved.Frontend == domain.FrontendNone {
				return nil
			}

			shouldInstall := installDeps
			if !cmd.Flags().Changed("install-deps") {
				shouldInstall, err = confirmInstallDeps()
				if err != nil {
					return err
				}
			}
			if !shouldInstall {
				return nil
			}
			pm, err := infra.ResolveNodePackageManager()
			if err != nil {
				return err
			}
			frontendDir := filepath.Join(resolved.OutputDir, "frontend")
			if err := infra.Run(frontendDir, pm, "install"); err != nil {
				return err
			}
			fmt.Println("Frontend dependencies installed.")
			return nil
		},
	}

	cmd.Flags().StringVar(&spec.ModulePath, "module", "", "Go module path, e.g. github.com/you/app")
	cmd.Flags().Var(&spec.Router, "router", "gin | chi | none")
	cmd.Flags().Var(&spec.Database, "db", "postgres | mysql | mongodb | redis | none")
	cmd.Flags().Var(&spec.Frontend, "frontend", "vite | nextjs | none")
	cmd.Flags().BoolVar(&spec.Force, "force", false, "overwrite an existing non-empty output directory")
	cmd.Flags().BoolVar(&spec.Docker, "docker", false, "generate a multi-stage Dockerfile and .dockerignore for the backend")
	cmd.Flags().BoolVar(&installDeps, "install-deps", false, "install frontend dependencies after scaffolding (skips the prompt)")
	return cmd
}
