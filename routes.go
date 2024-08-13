package main

import (
	"net/http"
)

//Create route model
type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

//Create data type "Routes" (slice)
type Routes []Route

//Various routes
var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoShow,
	},

	//Route to take in and store some JSON
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		TodoCreate,
	},
}