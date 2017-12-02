package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/go-pg/pg"
	"github.com/q231950/alethea/datastorage"
	"github.com/q231950/alethea/model"
)

func main() {
	log.SetHandler(cli.New(os.Stderr))
	log.SetLevel(log.InfoLevel)

	log.Info("Starting alethea...")

	ds := datastorage.NewDataStorage(database())
	ds.CreateIncidentsTable()

	http.HandleFunc("/post", postStatusHandler)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func database() pg.DB {
	user := os.Getenv("ALETHEA_POSTGRESQL_USER")
	password := os.Getenv("ALETHEA_POSTGRESQL_PASSWORD")
	database := os.Getenv("ALETHEA_POSTGRESQL_DATABASE")
	return *pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: database,
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Infof("Request: %s", r.Method)
}

func postStatusHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("Unable to read request's body."))
		return
	}

	if len(body) == 0 {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("The request body must have content."))
		return
	}
	ds := datastorage.NewDataStorage(database())
	incident := model.NewIncident()
	ds.LogIncident(incident)

	log.Infof("the body %s", body)
	w.WriteHeader(http.StatusOK)
}
