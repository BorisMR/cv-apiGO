package main

import "net/http"


type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

	Route{
		"Full",
		"GET",
		"/full",
		Full,
	},
	Route{
		"Contact",
		"GET",
		"/contact",
		Contact,
	},
	Route{
		"Laboral",
		"GET",
		"/laboral",
		Laboral,
	},
	Route{
		"Academic",
		"GET",
		"/academic",
		Academic,
	},
}