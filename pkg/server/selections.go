package server

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (s *Server) HandleMCRConnectionRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// TODO Some Error
		return
	}

	connections := make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&connections)
	if err != nil {
		panic(err)
	}

	connectedSources := make(map[string]string)
	for dest, source := range connections {
		if err := s.Panel.SourcesAndDestinations.MapSourceToDestination(uuid.MustParse(source), uuid.MustParse(dest)); err != nil {
			// TODO Error
		}
		connectedSources[dest] = source
	}

	w.Header().Add("Content-Type", "application/json")

	data, err := json.Marshal(connectedSources)
	if err != nil {
		// TODO error
	}
	w.Write(data)
}
