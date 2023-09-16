package main

import (
	"log"
	"net/http"

	"github.com/me-engi/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", nil)) // Use "localhost" instead of "Localhost" and correct ListenAndServe
}
