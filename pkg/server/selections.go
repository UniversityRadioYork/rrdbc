package server

import (
	"encoding/json"
	"net/http"

	"github.com/UniversityRadioYork/rrdbc/pkg/switcher"
)

func HandleMCRConnectionRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// TODO Some Error
		return
	}

	connections := make(map[string]switcher.Source)
	err := json.NewDecoder(r.Body).Decode(&connections)
	if err != nil {
		panic(err)
	}

	connectedSources := make(map[string]string)
	for dest, source := range connections {
		if err := switcher.MapSourceToDestination(source, dest); err != nil {
			// TODO Error
		}
		connectedSources[dest] = source.ShortName
	}

	w.Header().Add("Content-Type", "application/json")

	data, err := json.Marshal(connectedSources)
	if err != nil {
		// TODO error
	}
	w.Write(data)
}
