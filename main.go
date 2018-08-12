package main

import (
	"duplicator/api"
	"log"
	"net/http"
	"time"
)

func main() {

	srv := &http.Server{
		Addr: ":8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      api.NewAPI(),
	}

	log.Fatal(srv.ListenAndServe())
}
