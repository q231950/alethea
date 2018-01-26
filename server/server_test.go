package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	server := Server{}
	assert.NotNil(t, server)
}
