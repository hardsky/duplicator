package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-pg/pg"
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

type DuplicateResponse struct {
	Duplicate bool `json:"duplicate"`
}

func handleDuplicate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user1 := vars["userId1"]
	user2 := vars["userId2"]

	var count int
	db.QueryOne(pg.Scan(&count), `SELECT COUNT (1) FROM
(SELECT ip_addr FROM conn_log
WHERE user_id = ?) c1
INNER JOIN
(SELECT ip_addr FROM conn_log
        WHERE user_id = ?) c2
ON c1.ip_addr = c2.ip_addr
`, user1, user2)
	res := &DuplicateResponse{count >= 2}
	b, _ := json.Marshal(res)
	w.Write(b)
}

var db *pg.DB

func init() {
	db = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "jQnas3wed",
		Database: "duplicator",
	})
}
