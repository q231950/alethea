// Database

package database

import (
	"os"

	"github.com/go-pg/pg"
)

// Database
func PostgresqlDatabase() pg.DB {
	user := os.Getenv("ALETHEA_POSTGRESQL_USER")
	password := os.Getenv("ALETHEA_POSTGRESQL_PASSWORD")
	database := os.Getenv("ALETHEA_POSTGRESQL_DATABASE")
	return *pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: database,
	})
}
