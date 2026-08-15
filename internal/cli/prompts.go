package cli

import (
	"github.com/charmbracelet/huh"
	"github.com/sajanv88/gocli/internal/domain"
)

func confirmInstallDeps() (bool, error) {
	var install bool
	err := huh.NewConfirm().Title("Install frontend dependencies now?").
		Affirmative("Yes").
		Negative("No").
		Value(&install).
		Run()
	return install, err
}

func resolveSpec(spec domain.ProjectSpec) (domain.ProjectSpec, error) {
	var fields []huh.Field
	if spec.Name == "" {
		fields = append(fields, huh.NewInput().Title("Project name").Value(&spec.Name))
	}
	if spec.ModulePath == "" {
		fields = append(fields, huh.NewInput().Title("Module path").Value(&spec.ModulePath))
	}
	if spec.Router == "" {
		fields = append(fields, huh.NewSelect[domain.RouterOption]().Title("Router").
			Options(huh.NewOption("Gin", domain.RouterGin), huh.NewOption("Chi", domain.RouterChi)).
			Value(&spec.Router))
	}
	if spec.Database == "" {
		fields = append(fields, huh.NewSelect[domain.DatabaseOption]().Title("Database").
			Options(
				huh.NewOption("PostgreSQL", domain.DBPostgres), huh.NewOption("MySQL", domain.DBMySQL),
				huh.NewOption("MongoDB", domain.DBMongoDB), huh.NewOption("Redis", domain.DBRedis),
				huh.NewOption("None", domain.DBNone),
			).Value(&spec.Database))
	}
	if spec.Frontend == "" {
		fields = append(fields, huh.NewSelect[domain.FrontendOption]().Title("Frontend").
			Options(
				huh.NewOption("Vite (React)", domain.FrontendVite), huh.NewOption("Next.js", domain.FrontendNextJS),
				huh.NewOption("None", domain.FrontendNone),
			).Value(&spec.Frontend))
	}

	if len(fields) == 0 {
		return spec, nil // every flag was supplied. Fully non-interactive run
	}
	err := huh.NewForm(huh.NewGroup(fields...)).Run()
	return spec, err
}
