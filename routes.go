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
		"DbIndex",
		"GET",
		"/",
		DbIndex,
	},
	Route{
		"CollectionIndex",
		"GET",
		"/{dbName}",
		CollectionIndex,
	},
	Route{
		"QueryCollection",
		"GET",
		"/{dbName}/{collectionName}",
		QueryCollection,
	},
	Route{
		"InsertCollection",
		"POST",
		"/{dbName}/{collectionName}",
		InsertCollection,
	},
}
