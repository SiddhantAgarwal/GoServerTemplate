package main

import (
	"log"
	"net/http"

	"github.com/siddhant/GoServerTemplate/routes"
)

func main() {
	router := routes.NewRouter().StrictSlash(true)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
