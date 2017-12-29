package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = Logger(route.HandlerFunc, route.Name)
		router.Methods(route.Method).Path(route.Pattern).Handler(handler).Name(route.Name)
	}
	var notFound NotFound
	var notAllowed NotAllowed
	router.NotFoundHandler = notFound
	router.MethodNotAllowedHandler = notAllowed
	return router
}

func indexWebHandler(w http.ResponseWriter, r *http.Request) {
	// log.Println(r.URL.Query())
	// log.Println(mux.Vars(r))
	mapWelcome := map[string]bool{"isServerDown": false}
	mapped, _ := json.Marshal(mapWelcome)
	sendJson(w, mapped)
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		indexWebHandler,
	},
}

func sendJson(w http.ResponseWriter, resp []byte) {
	w.Header().Set("Content-type", "application/json")
	w.Write(resp)
}
