package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Server struct {
	Port       int
	MetaGroups [16]struct {
		GrpName string
		Members [16]struct {
			ShortName string
			LongName  string
		}
	}
}

func (s *Server) Start() {

	http.Handle("/control/", http.StripPrefix("/control/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/control/meta", func(w http.ResponseWriter, r *http.Request) {
		json, err := json.Marshal(s.MetaGroups)
		if err != nil {
			// TODO
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(json)
	})

	http.ListenAndServe(fmt.Sprintf(":%v", s.Port), nil)

}
