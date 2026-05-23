package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/SiddhantAgarwal/GoServerTemplate/internal/handler"
)

type NotFound struct{}

func (NotFound) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write([]byte(`{"error": "Not Found"}`))
}

type NotAllowed struct{}

func (NotAllowed) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	_, _ = w.Write([]byte(`{"error": "Not Allowed"}`))
}

func NewRouter(health *handler.HealthHandler) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", health.HandleIndex).Methods(http.MethodGet).Name("Index")
	router.HandleFunc("/health", health.HandleHealth).Methods(http.MethodGet).Name("Health")

	router.NotFoundHandler = NotFound{}
	router.MethodNotAllowedHandler = NotAllowed{}

	return router
}
