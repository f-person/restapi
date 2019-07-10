package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Book struct {
	Id     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// get all books
func getBooks(w http.ResponseWriter, r *http.Request) {

}

// get single book
func getBook(w http.ResponseWriter, r *http.Request) {

}

// create a new book
func createBook(w http.ResponseWriter, r *http.Request) {

}

// udate a book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// init router
	r := mux.NewRouter()

	// route handlers / endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
