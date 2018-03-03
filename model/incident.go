package model

import (
	"encoding/json"
	"errors"

	"github.com/q231950/alethea/ci"
)

// Incident describes an incident that occurred
type Incident struct {
	CI          string
	Identifier  string
	Failed      bool
	Committer   string
	Project     string
	BuildNumber string
	BuildUrl    string
}

// NewIncident creates an incident with a random identifier
func NewIncidentFromJson(c ci.CI, jsonblob []byte) (Incident, error) {
	if c == ci.Circle {
		var incident *ci.CircleCIIncident
		error := json.Unmarshal(jsonblob, &incident)
		return Incident{incident.CI(),
			incident.Identifier(),
			incident.Failed(),
			incident.Committer(),
			incident.Project(),
			incident.BuildNumber(),
			incident.BuildUrl()}, error
	}
	return Incident{"", "", true, "", "", "", ""},
		errors.New("Unable to unmarshal post body to incident")
}

func (i *Incident) String() string {
	return i.CI
}
