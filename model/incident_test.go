package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIncident(t *testing.T) {
	incident, err := NewIncident()
	assert.Nil(t, err)
	assert.NotNil(t, incident.Identifier, "An incident's identifier should never be nil")
	assert.Equal(t, incident.Source, "source")
	assert.Equal(t, incident.Value, "value")
}
