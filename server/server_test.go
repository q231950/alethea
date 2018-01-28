package server

import (
	"errors"
	"net/http"
	"testing"

	"github.com/q231950/alethea/model"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	server := Server{}
	assert.NotNil(t, server)
}

func TestPostStatusHandler(t *testing.T) {
	server := Server{}
	w := *new(http.ResponseWriter)
	err := errors.New("some error when creating the build result")
	incident := model.Incident{}
	server.handleBuildResult(incident, err, w)
	assert.Equal(t, w.Header, http.StatusInternalServerError)
}
