package api

import (
	"net/http"

	"github.com/hardsky/duplicator/db"
	"github.com/hardsky/duplicator/redis"

	"github.com/gorilla/mux"
)

// Opts contains db and cache.
type Opts struct {
	Db    *db.DB
	Cache *redis.Cache
}

// NewAPI constructs new api
func NewAPI(opt *Opts) *API {
	router := mux.NewRouter()
	res := &API{
		d: opt.Db,
		c: opt.Cache,
		h: router,
	}

	//service prefix
	s := router.PathPrefix("/duplicator/").Subrouter()

	//service routes
	s.HandleFunc("/duplicate/{userId1:[0-9]+}/{userId2:[0-9]+}", res.handleDuplicate).Methods("GET")

	return res
}

//API contains methods that implements routes.
type API struct {
	d *db.DB
	c *redis.Cache
	h http.Handler
}

// Handler returns http.Handler with api routes.
func (p *API) Handler() http.Handler {
	return p.h
}
