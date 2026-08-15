package db

import (
	"github.com/sajanv88/gocli/internal/adapter/db/mongo"
	"github.com/sajanv88/gocli/internal/adapter/db/mysql"
	"github.com/sajanv88/gocli/internal/adapter/db/postgres"
	"github.com/sajanv88/gocli/internal/adapter/db/redis"
	"github.com/sajanv88/gocli/internal/domain"
)

func All() map[domain.DatabaseOption]domain.DBGenerator {
	return map[domain.DatabaseOption]domain.DBGenerator{
		domain.DBPostgres: postgres.Generator{},
		domain.DBMySQL:    mysql.Generator{},
		domain.DBMongoDB:  mongo.Generator{},
		domain.DBRedis:    redis.Generator{},
	}
}
