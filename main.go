package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	//s := r.PathPrefix("/duplicator/").Subrouter()
	r.HandleFunc("/duplicate/{userId1:[0-9]+}/{userId2:[0-9]+}", handleDuplicate).Methods("GET")

	srv := &http.Server{
		Addr: ":8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	srv.ListenAndServe()
}

func handleDuplicate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "userId1: %v\n", vars["userId1"])
	fmt.Fprintf(w, "userId2: %v\n", vars["userId2"])
}
