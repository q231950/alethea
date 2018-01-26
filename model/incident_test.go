package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testNewIncident(t *testing.T) {
	incident := NewIncident()
	assert.NotNil(t, incident.Identifier, "An incident's identifier should never be nil")
	assert.Equal(t, incident.Source, "source")
	assert.Equal(t, incident.Value, "value")
}
