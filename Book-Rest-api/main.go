package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Model for a book

type Book struct {
	ID     string `json:"id"`
	ISBN   string `json:"isbn"`
	Title  string `json:"title"`
	Author Author `json:"author"`
}

//author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//init slice of books

var books []Book

func main() {

	fmt.Println("Book API started...")

	//router using mux
	r := mux.NewRouter()

	books = append(books, Book{
		ID:    "1",
		ISBN:  "55677",
		Title: "Book one",
		Author: Author{
			Firstname: "Rajat",
			Lastname:  "Singh",
		},
	})
	books = append(books, Book{
		ID:    "2",
		ISBN:  "44677",
		Title: "Book Two",
		Author: Author{
			Firstname: "Lokesh",
			Lastname:  "Singh",
		},
	})
	books = append(books, Book{
		ID:    "3",
		ISBN:  "44688",
		Title: "Book Three",
		Author: Author{
			Firstname: "Prachi",
			Lastname:  "Singh",
		},
	})

	//endpoints of api
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	//starting server
	log.Fatal(http.ListenAndServe(":8080", r))

}

// get books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}

// get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	idint, err := strconv.Atoi(id)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid ID")
	}
	if idint > len(books) {
		json.NewEncoder(w).Encode("Invalid ID")
		return
	}
	for _, val := range books {
		if val.ID == id {
			json.NewEncoder(w).Encode(val)
		}
	}

}

// create a book
func createBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var newBook Book
	json.NewDecoder(r.Body).Decode(&newBook)
	books = append(books, newBook)
	json.NewEncoder(w).Encode(newBook)
	json.NewEncoder(w).Encode("book addded successfully")

}

//update a book
func updateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}

}

//delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
