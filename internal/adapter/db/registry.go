package db

import "github.com/sajanv88/gocli/internal/domain"

func All() map[domain.DatabaseOption]domain.DBGenerator {
	return map[domain.DatabaseOption]domain.DBGenerator{
		"postgres": PostgresGenerator{},
		"mysql":    MySQLGenerator{},
		"mongodb":  MongoGenerator{},
		"redis":    RedisGenerator{},
	}
}
