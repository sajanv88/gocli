package domain

import (
	"errors"
	"fmt"
)

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

func (r *RouterOption) String() string { return string(*r) }

func (r *RouterOption) Set(s string) error {
	switch RouterOption(s) {
	case RouterGin, RouterChi, RouterNone:
		*r = RouterOption(s)
		return nil
	default:
		return fmt.Errorf("invalid router %q: must be one of gin, chi, none", s)
	}
}

func (r *RouterOption) Type() string { return "router" }

func (d *DatabaseOption) String() string { return string(*d) }

func (d *DatabaseOption) Set(s string) error {
	switch DatabaseOption(s) {
	case DBPostgres, DBMySQL, DBMongoDB, DBRedis, DBNone:
		*d = DatabaseOption(s)
		return nil
	default:
		return fmt.Errorf("invalid database %q: must be one of postgres, mysql, mongodb, redis, none", s)
	}
}

func (d *DatabaseOption) Type() string { return "database" }

func (f *FrontendOption) String() string { return string(*f) }

func (f *FrontendOption) Set(s string) error {
	switch FrontendOption(s) {
	case FrontendVite, FrontendNextJS, FrontendNone:
		*f = FrontendOption(s)
		return nil
	default:
		return fmt.Errorf("invalid frontend %q: must be one of vite, nextjs, none", s)
	}
}

func (f *FrontendOption) Type() string { return "frontend" }

func (s ProjectSpec) Validate() error {

	if s.Name == "" {
		return errors.New("name is required")
	}

	if s.ModulePath == "" {
		return errors.New("modulePath is required (e.g. github.com/your_username/project)")
	}

	return nil
}
