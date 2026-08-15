package router

import "github.com/sajanv88/gocli/internal/domain"

func All() map[domain.RouterOption]domain.RouterGenerator {
	return map[domain.RouterOption]domain.RouterGenerator{
		"gin": GinGenerator{},
		"chi": ChiGenerator{},
	}
}
