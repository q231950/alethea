package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	server := Server{}
	assert.NotNil(t, server)
}

func TestPostStatusHandler(t *testing.T) {
	// server := Server{}
	// w := new(http.ResponseWriter)
	// r := new(http.Request)
	// r.Method = ""
	// err := Error{"some error when creating the build result"}
	// server.handleBuildResult(nil, err, w)
	// assert.Equal(t, w.Header, )
}
