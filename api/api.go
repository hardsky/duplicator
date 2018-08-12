package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewAPI() http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix("/duplicator/").Subrouter()
	s.HandleFunc("/duplicate/{userId1:[0-9]+}/{userId2:[0-9]+}", handleDuplicate).Methods("GET")

	return s
}
