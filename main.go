package main

import (
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/q231950/alethea/database"
	"github.com/q231950/alethea/datastorage"
)

func main() {
	log.SetHandler(cli.New(os.Stderr))
	log.SetLevel(log.InfoLevel)

	log.Info("Starting alethea...")

	database := database.PostgresqlDatabase()
	dataStorage := datastorage.NewDataStorage(database)
	dataStorage.CreateIncidentsTable()

	server := Server.New(dataStorage)
}
