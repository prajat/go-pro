package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"fmt"
)

//User is a struct
type User struct {
	FullName string `json:"fullName"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

//Post is a struct
type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author User   `json:"author"`
}

var posts []Post = []Post{}

func main() {

	// use mux library for the routing part
	fmt.Println("starting the magicpin service")
	router := mux.NewRouter()
	router.HandleFunc("/posts", addItem).Methods("POST")
	router.HandleFunc("/posts", getAllPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	log.Fatal(http.ListenAndServe("localhost:5000", router))

}

func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "appliaction/json")

	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("error"))
		return
	}
	if id >= len(posts) { //error checking
		w.WriteHeader(404)
		w.Write([]byte("no post found with specified id"))
		return
	}
	var updatedPost Post
	json.NewDecoder(r.Body).Decode(&updatedPost)
	posts[id] = updatedPost
	json.NewEncoder(w).Encode(updatedPost)

}

func getPost(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("id cant be converted to integer"))
		return
	}
	if id >= len(posts) { //error checking
		w.WriteHeader(404)
		w.Write([]byte("no post found with specified id"))
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(posts[id])

}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(posts)

}

func addItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost)
	posts = append(posts, newPost)

	json.NewEncoder(w).Encode(posts)
}
