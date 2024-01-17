package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ssr0016/webapp/middleware"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", homeGetHandler).Methods("GET")
	r.HandleFunc("/", homePostHandler).Methods("POST")
	r.HandleFunc("/register", registerGetHandler).Methods("GET")
	r.HandleFunc("/register", registerPostHandler).Methods("POST")
	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")
	r.HandleFunc("/admin", middleware.AuthRequired(adminGetHandler)).Methods("GET")

	fileServer := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))
	return r
}
