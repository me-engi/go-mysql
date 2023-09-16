package routes

import (
	"github.com/gorilla/mux"
	"github.com/me-engi/go-bookstore/pkg/controllers"
)

func RegisterBookStoreRoutes(router *mux.Router) {
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books", controllers.GetBook).Methods("GET")

	router.HandleFunc("/books/{bookId}", controllers.GetBookId).Methods("GET")
	router.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
