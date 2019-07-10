package main

import (
	//"encoding/json"
	"github.com/gorilla/mux"
	//"log"
	//"math/rand"
	//"net/http"
	//"strconv"
)

func main() {
	// init router
	r := mux.NewRouter()

	// route handlers / endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

}
