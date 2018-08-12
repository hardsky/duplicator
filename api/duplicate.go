package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type duplicateResponse struct {
	Duplicate bool `json:"duplicate"`
}

// Implement /duplicate/{userId1}/{userId2}
// Determines when two users are duplicates.
func (p *API) handleDuplicate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user1 := vars["userId1"]
	user2 := vars["userId2"]

	ctxLog := log.WithFields(log.Fields{
		"userID1": user1,
		"userID2": user2,
	})

	if p.c.IsDuplicate(user1, user2) {
		ctxLog.Debug("cache matched")
		writeTrue(w)
		return
	}
	ctxLog.Debug("cache missed")

	res := &duplicateResponse{p.d.IsDuplicate(user1, user2)}
	if res.Duplicate {
		p.c.Duplicate(user1, user2)
	}

	b, _ := json.Marshal(res)
	w.Write(b)
}

func writeTrue(w http.ResponseWriter) {
	b, _ := json.Marshal(&duplicateResponse{true})
	w.Write(b)
}
