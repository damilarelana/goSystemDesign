package middleware

import (
	mux "github.com/gorilla/mux"
)

// NewRouter initializes the router
// - setups up all known routes
// - returns the router
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true) // initialize a new mux router

	for _, route := range routes { // setup router with all the defined `routes` in routes
		router.Methods(route.Method).Path(route.Path).Name(route.Keyword).Handler(route.HandlerFunc)
	}

	return router
}
