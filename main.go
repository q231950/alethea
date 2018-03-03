package main

import (
	"flag"
	"os"
	"strconv"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/q231950/alethea/datastorage"
	"github.com/q231950/alethea/model"
	"github.com/q231950/alethea/server"
)

const defaultPort int = 8080

func main() {
	log.SetHandler(cli.New(os.Stderr))
	log.SetLevel(log.InfoLevel)

	log.Info("Starting alethea...")

	dataStorage := datastorage.New()
	incidentModel := new(model.Incident)
	dataStorage.CreateTable(incidentModel)

	var port string
	var portEnvironmentVariable = os.Getenv("PORT")
	if len(portEnvironmentVariable) > 0 {
		port = portEnvironmentVariable
	} else {
		p := flag.Int("port", defaultPort, "help message for flagname")
		flag.Parse()
		port = strconv.Itoa(*p)
	}

	log.Infof("Creating new alethea server, serving port %s", port)
	server := server.NewServer(dataStorage, port)
	server.Serve()
}
