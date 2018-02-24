package main

import (
	"flag"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/q231950/alethea/datastorage"
	"github.com/q231950/alethea/server"
)

const defaultPort = 8080

func main() {
	log.SetHandler(cli.New(os.Stderr))
	log.SetLevel(log.InfoLevel)

	log.Info("Starting alethea...")

	dataStorage := datastorage.New()
	dataStorage.CreateIncidentsTable()

	p := flag.Int("port", defaultPort, "help message for flagname")
	flag.Parse()

	server := server.NewServer(dataStorage, *p)
	server.Serve()
}
