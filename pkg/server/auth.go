package server

import (
	"net/http"
)

type authHandler struct {
	Next          http.Handler
	AdminRequired bool
	Users         map[string]struct {
		Password string
		Admin    bool
	}
}

func authFail(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(w, "Unauthorised", http.StatusUnauthorized)
}

func (a *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	u, p, ok := r.BasicAuth()
	if !ok {
		authFail(w)
		return
	}

	if _, ok := a.Users[u]; !ok {
		authFail(w)
		return
	}

	if p != a.Users[u].Password {
		authFail(w)
		return
	}

	if a.AdminRequired && !a.Users[u].Admin {
		authFail(w)
		return
	}

	a.Next.ServeHTTP(w, r)
}
