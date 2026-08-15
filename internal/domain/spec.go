package domain

import "errors"

type RouterOption string
type DatabaseOption string
type FrontendOption string

const (
	RouterGin  RouterOption = "gin"
	RouterChi  RouterOption = "chi"
	RouterNone RouterOption = "none"
)

const (
	DBPostgres DatabaseOption = "postgres"
	DBMySQL    DatabaseOption = "mysql"
	DBMongoDB  DatabaseOption = "mongodb"
	DBRedis    DatabaseOption = "redis"
	DBNone     DatabaseOption = "none"
)

const (
	FrontendVite   FrontendOption = "vite"
	FrontendNextJS FrontendOption = "nextjs"
	FrontendNone   FrontendOption = "none"
)

type ProjectSpec struct {
	Name       string
	ModulePath string
	Router     RouterOption
	Database   DatabaseOption
	Frontend   FrontendOption
	OutputDir  string
}

func (s ProjectSpec) Validate() error {

	if s.Name == "" {
		return errors.New("name is required")
	}

	if s.ModulePath == "" {
		return errors.New("modulePath is required (e.g. github.com/your_username/project)")
	}

	return nil
}
