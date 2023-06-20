package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Q map[string]string

type handlerFunc func(http.ResponseWriter, *http.Request)

func RegisterHandler(
	router *mux.Router, // parent Router
	method string, // HTTP method
	path string, // URL path
	query Q, // URL query parameters (can be nil)
	handler handlerFunc, // handler function
) *mux.Route {
	// new route for the path.
	route := router.Path(path)

	// specify the method unless it's an empty string or "*".
	if method != "" && method != "*" {
		route = route.Methods(method)
	}

	// specify the query params (if any)
	for key, value := range query {
		route.Queries(key, value)
	}

	// setup the handler which gets us a new route
	if handler != nil {
		route.HandlerFunc(handler)
	}

	return route
}
