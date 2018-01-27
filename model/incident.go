package model

import uuid "github.com/satori/go.uuid"

// Incident describes an incident that occurred
type Incident struct {
	Identifier string
	Source     string
	Value      string
}

// NewIncident creates an incident with a random identifier
func NewIncident() (Incident, error) {
	identifier, err := uuid.NewV4()
	incident := Incident{Identifier: identifier.String(), Source: "source", Value: "value"}
	return incident, err
}
