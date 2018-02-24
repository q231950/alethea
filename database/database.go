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
	address := os.Getenv("ALETHEA_POSTGRESQL_ADDRESS") + ":" + os.Getenv("ALETHEA_POSTGRESQL_PORT")
	return *pg.Connect(&pg.Options{
		Addr:     address,
		User:     user,
		Password: password,
		Database: database,
	})
}
