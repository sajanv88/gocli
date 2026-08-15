package router

import "github.com/sajanv88/gocli/internal/domain"

func All() map[domain.RouterOption]domain.RouterGenerator {
	return map[domain.RouterOption]domain.RouterGenerator{
		domain.RouterGin: GinGenerator{},
		domain.RouterChi: ChiGenerator{},
	}
}
