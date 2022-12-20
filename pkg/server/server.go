package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/UniversityRadioYork/rrdbc/pkg/metadata"
	"github.com/UniversityRadioYork/rrdbc/pkg/panel"
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
	Users map[string]struct {
		Password string
		Admin    bool
	}
	Panel panel.MCRPanel
}

func (s *Server) Start() {

	http.HandleFunc("/meta", func(w http.ResponseWriter, r *http.Request) {
		// TODO - cache
		data, err := json.Marshal(metadata.GetStreamMetadata(s.Panel.SourcesAndDestinations.Destinations))
		if err != nil {
			// TODO Error
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(data)
	})

	http.Handle("/control/",
		&authHandler{
			Next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				s.Panel.RenderTemplate(w, "templates/index.html")
			}),
			Users: s.Users,
		},
	)

	http.Handle("/control/main.js",
		&authHandler{
			Next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				s.Panel.RenderTemplate(w, "templates/main.js")
			}),
			Users: s.Users,
		},
	)

	http.Handle("/control/styles.css",
		&authHandler{
			Next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.ServeFile(w, r, "static/styles.css")
			}),
			Users: s.Users,
		},
	)

	http.Handle("/control/meta", &authHandler{
		Next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			json, err := json.Marshal(s.MetaGroups)
			if err != nil {
				// TODO
			}
			w.Header().Add("Content-Type", "application/json")
			w.Write(json)
		}),
		Users: s.Users,
	})

	http.Handle("/control/user", &authHandler{
		Next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			u, _, _ := r.BasicAuth()
			fmt.Fprint(w, u)
		}),
		Users: s.Users,
	})

	http.Handle("/control/take", &authHandler{
		Next:  http.HandlerFunc(s.HandleMCRConnectionRequest),
		Users: s.Users,
	})

	http.ListenAndServe(fmt.Sprintf(":%v", s.Port), nil)

}
