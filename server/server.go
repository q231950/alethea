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
	dataStorage datastorage.DataStorage
}

// NewServer returns an instance of Server
func NewServer(ds datastorage.DataStorage) Server {
	server := Server{dataStorage: ds}
	http.HandleFunc("/post", server.postStatusHandler)
	http.HandleFunc("/", server.handler)
	return server
}

// Serve starts serving the service
func (server *Server) Serve() error {
	return http.ListenAndServe(":443", nil)
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

	buildResult, err := model.NewIncident()
	server.handleBuildResult(buildResult, err, w)

	log.Infof("the body %s", body)
}

func (server *Server) handleBuildResult(buildResult model.Incident, err error, w http.ResponseWriter) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Infof("Failed to create new build result from payload %s", err)
	} else {
		server.dataStorage.StoreIncident(buildResult)
		w.WriteHeader(http.StatusOK)
	}
}
