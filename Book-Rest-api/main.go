package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Model

type Book struct {
	ID     string `json:"id"`
	ISBN   string `json:"isbn"`
	Title  string `json:"title"`
	Author Author `json:"author"`
}

func main() {

	fmt.Println("Book API started...")

	//router using mux
	r := mux.NewRouter()

	//endpoints of api
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", upadteBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	//starting server
	log.Fatal(http.ListenAndServe(":8080", r))

}
