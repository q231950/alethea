package model

import uuid "github.com/satori/go.uuid"

// Incident describes an incident that occurred
type Incident struct {
	Identifier string
	Source     string
	Value      string
}

// New creates an incident with a random identifier
func NewIncident() Incident {
	identifier := uuid.NewV4().String()
	incident := Incident{Identifier: identifier, Source: "source", Value: "value"}
	return incident
}
