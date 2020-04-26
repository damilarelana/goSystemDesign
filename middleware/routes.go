package middleware

import (
	"net/http"
)

// Route defines a composite variable structure that maps https routes to a callable abbreviated name (keyword):
// - keyword abbreviation (to call it)
// - route's method
// - route's pattern (i.e. actual path)
// - route's handlerfunction
type Route struct {
	Keyword     string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

// Routes defines the actual variable type
type Routes []Route

// Initialize a typical routes [right now it is just one instance of that routes, additional routes would just use `Route()` again]
var routes = Routes{
	Route{
		"getAccount",            // keyword
		"GET",                   // method
		"/accounts/{accountID}", // actual path
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8") // set the request headers metadata
			w.Write([]byte("{\"result\":\"OK\"}"))
		},
	},
}
