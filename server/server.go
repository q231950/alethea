// Server

package server

import (
	"io/ioutil"
	"net/http"
	"os"

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
	return http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

// func (server *Server) Shutdown() error {
// 	http.DefaultServeMux.Server.Shutdown()
// }

func (server *Server) handler(w http.ResponseWriter, r *http.Request) {
	log.Infof("Request: %s", r.Method)
}

func (server *Server) postStatusHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("Endpoint `status` only accepts http `POST`."))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("Unable to read request's body."))
		return
	}

	log.Infof("Handling message to status with body: %s", body)

	if len(body) == 0 {
		w.Write([]byte("The request body must have content. It must not be empty."))
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	buildResult, err := model.NewIncident()
	server.handleBuildResult(buildResult, err, w)
}

func (server *Server) handleBuildResult(buildResult model.Incident, err error, w http.ResponseWriter) {
	if err != nil {
		log.Infof("Failed to create new build result from payload %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Infof("Storing status entry %s", buildResult)
	server.dataStorage.StoreIncident(buildResult)
	w.WriteHeader(http.StatusAccepted)
}
