package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, r := range routes {
		router.Name(r.Name).Methods(r.Method).Path(r.Pattern).Handler(r.HandleFunc)
	}
	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"MoviesList",
		"GET",
		"/movies",
		MoviesList,
	},
	Route{
		"MovieById",
		"GET",
		"/movies/{id}",
		MovieById,
	},
}
