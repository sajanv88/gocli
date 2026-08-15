package db

import "github.com/sajanv88/gocli/internal/domain"

func All() map[domain.DatabaseOption]domain.DBGenerator {
	return map[domain.DatabaseOption]domain.DBGenerator{
		domain.DBPostgres: PostgresGenerator{},
		domain.DBMySQL:    MySQLGenerator{},
		domain.DBMongoDB:  MongoGenerator{},
		domain.DBRedis:    RedisGenerator{},
	}
}
