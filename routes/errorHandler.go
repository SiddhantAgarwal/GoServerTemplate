package routes

import "net/http"

type NotFound func(w http.ResponseWriter, req *http.Request)

func (NotFound) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(404)
	w.Write([]byte(`{"Error": "Not Found"}`))
}

type NotAllowed func(w http.ResponseWriter, req *http.Request)

func (NotAllowed) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(405)
	w.Write([]byte(`{"Error": "Not Allowed"}`))
}
