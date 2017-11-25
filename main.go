package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
)

func main() {
	log.SetHandler(cli.New(os.Stderr))
	log.SetLevel(log.InfoLevel)

	log.Info("Starting alethea...")

	http.HandleFunc("/post", postStatusHandler)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
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

	log.Infof("the body %s", body)

	w.WriteHeader(http.StatusOK)
}
