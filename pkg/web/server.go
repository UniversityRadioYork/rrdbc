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
	Users map[string]struct {
		Password string
		Admin    bool
	}
}

func (s *Server) Start() {

	http.Handle("/control/",
		&AuthHandler{
			Next:  http.StripPrefix("/control/", http.FileServer(http.Dir("./static"))),
			Users: s.Users,
		},
	)

	http.Handle("/control/meta", &AuthHandler{
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

	http.Handle("/control/user", &AuthHandler{
		Next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			u, _, _ := r.BasicAuth()
			fmt.Fprint(w, u)
		}),
		Users: s.Users,
	})

	http.ListenAndServe(fmt.Sprintf(":%v", s.Port), nil)

}
