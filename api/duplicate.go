package api

import (
	"duplicator/db"
	"duplicator/redis"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type duplicateResponse struct {
	Duplicate bool `json:"duplicate"`
}

func handleDuplicate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user1 := vars["userId1"]
	user2 := vars["userId2"]

	if redis.IsDuplicate(user1, user2) {
		fmt.Println("cache matched")
		writeTrue(w)
		return
	}
	fmt.Println("cache missed")

	res := &duplicateResponse{db.IsDuplicate(user1, user2)}
	if res.Duplicate {
		redis.Duplicate(user1, user2)
	}

	b, _ := json.Marshal(res)
	w.Write(b)
}

func writeTrue(w http.ResponseWriter) {
	b, _ := json.Marshal(&duplicateResponse{true})
	w.Write(b)
}
