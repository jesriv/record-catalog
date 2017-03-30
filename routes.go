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
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"ReleaseIndex",
		"GET",
		"/releases",
		ReleaseIndex,
	},
	Route{
		"ReleaseIndex",
		"POST",
		"/releases",
		ReleaseCreate,
	},
	Route{
		"ReleaseShow",
		"GET",
		"/releases/{releaseId}",
		ReleaseShow,
	},
	Route{
		"ReleaseUpdate",
		"PUT",
		"/releases/{releaseId}",
		ReleaseUpdate,
	},
	Route{
		"Authenticate",
		"POST",
		"/auth",
		Authenticate,		
	},
}