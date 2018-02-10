// Server

package server

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/apex/log"
	"github.com/gorilla/mux"
	"github.com/q231950/alethea/datastorage"
	"github.com/q231950/alethea/model"
)

// Server serves the http API endpoint
type Server struct {
	dataStorage datastorage.DataStorage
	httpServer  http.Server
}

// NewServer returns an instance of Server
func NewServer(ds datastorage.DataStorage) Server {
	r := mux.NewRouter()
	httpServer := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: r,
	}

	server := Server{dataStorage: ds, httpServer: httpServer}

	r.HandleFunc("/post", server.postStatusHandler)
	r.HandleFunc("/fun", Fun)
	r.HandleFunc("/", server.handler)

	return server
}

// Serve starts serving the service
func (server *Server) Serve() error {
	return server.httpServer.ListenAndServe()
}

func Fun(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "fun")
	w.WriteHeader(http.StatusOK)
}

func (server *Server) handler(w http.ResponseWriter, r *http.Request) {
	log.Infof("Request: %s", r.Method)
	w.WriteHeader(http.StatusNotVeryOK)
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

	log.Infof("Handling message to status with body: %s %d", body, len(body))

	if len(body) == 0 {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("The request body must have content. It must not be empty."))
		return
	}

	buildResult, err := model.NewIncident()
	server.handleBuildResult(buildResult, err, w)
}

func (server *Server) handleBuildResult(buildResult model.Incident, err error, w http.ResponseWriter) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Infof("Failed to create new build result from payload %s", err)
		return
	}

	log.Infof("Storing status entry %s", buildResult)
	server.dataStorage.StoreIncident(buildResult)
	w.WriteHeader(http.StatusAccepted)
}
