package main

// import packages
import(
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

// structure (model)
type Book struct {
	ID	string `json:"id"`
	Isbn	string `json:"isbn"`
	Title	string `json:"title"`
	Author	*Author `json:"author"`
}

// author structure
type Author struct {
	FirstName	string `json:firstname`
	LastName	string `json:lastname`
}

// get all books
func getBooks(w http.ResponseWriter, r *http.Request) {

}

// get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	
}

// create new book
func createBook(w http.ResponseWriter, r *http.Request) {
	
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

	// route handler & endpoint
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// run server
	log.Fatal(http.ListenAndServe(":8000", router))
	
}

