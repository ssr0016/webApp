package middleware

import (
	"net/http"

	"github.com/ssr0016/webapp/sessions"
)

func AuthRequired(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := sessions.Store.Get(r, "session")
		_, ok := session.Values["USERID"]
		if !ok {
			http.Redirect(w, r, "/login", 302)
			return
		}

		handler.ServeHTTP(w, r)
	}
}
