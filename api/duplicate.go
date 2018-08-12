package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type duplicateResponse struct {
	Duplicate bool `json:"duplicate"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

// Implement /duplicate/{userId1}/{userId2}
// Determines when two users are duplicates.
func (p *API) handleDuplicate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ctxLog := log.WithFields(log.Fields{
		"userID1": vars["userId1"],
		"userID2": vars["userId2"],
	})

	userID1, err := strconv.ParseInt(vars["userId1"], 10, 64)
	if err != nil || userID1 <= 0 {
		ctxLog.WithError(err).Error("wrong user Id")
		writeResponse(w, &ErrorResponse{"wrong user Id"}, http.StatusBadRequest)
		return
	}
	userID2, err := strconv.ParseInt(vars["userId2"], 10, 64)
	if err != nil || userID2 <= 0 {
		ctxLog.WithError(err).Error("wrong user Id")
		writeResponse(w, &ErrorResponse{"wrong user Id"}, http.StatusBadRequest)
		return
	}

	// check in cache for same previous request, that returned true
	// if we found that some user ids are duplicate, that are not changed
	if p.c.IsDuplicate(userID1, userID2) {
		ctxLog.Debug("cache matched")
		writeResponse(w, &duplicateResponse{true}, http.StatusOK)
		return
	}
	ctxLog.Debug("cache missed")

	// run query in db to determine duplicates or not
	res := &duplicateResponse{p.d.IsDuplicate(userID1, userID2)}
	if res.Duplicate {
		p.c.Duplicate(userID1, userID2)
	}

	writeResponse(w, res, http.StatusOK)
}

func writeResponse(w http.ResponseWriter, body interface{}, status int) {
	b, _ := json.Marshal(body)
	w.WriteHeader(status)
	w.Write(b)
	w.Header().Set("Content-Type", "application/json")
}
