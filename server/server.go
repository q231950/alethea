// Server

package server

import (
	"io/ioutil"
	"net/http"

	"github.com/apex/log"
	"github.com/q231950/alethea/datastorage"
	"github.com/q231950/alethea/model"
)

// Server serves the http API endpoint
type Server struct {
	dataStorage *datastorage.DataStorage
}

// NewServer returns an instance of Server
func NewServer(ds *datastorage.DataStorage) Server {
	server := Server{dataStorage: ds}
	http.HandleFunc("/post", server.postStatusHandler)
	http.HandleFunc("/", server.handler)
	http.ListenAndServe(":8080", nil)
	return server
}

func (server *Server) handler(w http.ResponseWriter, r *http.Request) {
	log.Infof("Request: %s", r.Method)
}

func (server *Server) postStatusHandler(w http.ResponseWriter, r *http.Request) {
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

	incident := model.NewIncident()
	server.dataStorage.StoreIncident(incident)

	log.Infof("the body %s", body)
	w.WriteHeader(http.StatusOK)
}
