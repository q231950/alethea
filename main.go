package main

import (
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/q231950/alethea/datastorage"
	"github.com/q231950/alethea/server"
)

func main() {
	log.SetHandler(cli.New(os.Stderr))
	log.SetLevel(log.InfoLevel)

	log.Info("Starting alethea...")

	dataStorage := datastorage.New()
	dataStorage.CreateIncidentsTable()

	server := server.NewServer(dataStorage, os.Getenv("PORT"))
	server.Serve()
}
