package main

import (
	"fmt"
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
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func postStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
