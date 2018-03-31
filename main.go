package main

// import packages
import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// structure (model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// author structure
type Author struct {
	Firstname string `json:firstname`
	Lastname  string `json:lastname`
}

// init books as a slice Book struct
var books []Book

// get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get params
	params := mux.Vars(r)

	// loop through and find book id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// create new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	// set a random id for fun, dont do in production
	book.ID = strconv.Itoa(rand.Intn(10000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// update a book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

// delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// initialize router
	router := mux.NewRouter()

	// mock data before implementing DB
	books = append(books, Book{ID: "1", Isbn: "525232", Title: "A cool book", Author: &Author{Firstname: "Jared", Lastname: "Dun"}})

	books = append(books, Book{ID: "2", Isbn: "234232", Title: "An awsome book", Author: &Author{Firstname: "Erlic", Lastname: "Bachman"}})

	// route handler & endpoint
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// run server
	log.Fatal(http.ListenAndServe(":8000", router))
}
