package router

import (
	"github.com/sajanv88/gocli/internal/adapter/router/chi"
	"github.com/sajanv88/gocli/internal/adapter/router/gin"
	"github.com/sajanv88/gocli/internal/domain"
)

func All() map[domain.RouterOption]domain.RouterGenerator {
	return map[domain.RouterOption]domain.RouterGenerator{
		domain.RouterGin: gin.Generator{},
		domain.RouterChi: chi.Generator{},
	}
}
