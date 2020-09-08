package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

//struct for the table user
type User struct {
	ID        gocql.UUID `json:"id"`
	FirstName string     `json:"firstname"`
	LastName  string     `json:"lastname"`
	Email     string     `json:"email"`
	Age       int        `json:"age"`
	City      string     `json:"city"`
}

var session *gocql.Session

func main() {
	fmt.Println("Starting API...")

	//DB connection

	var err error
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "streamdemoapi"
	session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("database connected successfully")

	//Routes
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", heartbeat)
	router.HandleFunc("/newUser", newUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))

}
func newUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	if err := session.Query(`INSERT INTO users (id,age,city,firstname,lastname) VALUES (?,?,?,?,?)`, gocql.TimeUUID(), user.Age, user.City, user.FirstName, user.LastName).Exec(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("a user is added to DB")
	json.NewEncoder(w).Encode("you added the user succesully")

}

//struct for heartbeat response
type heartbeatResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	hb := heartbeatResponse{Status: "OK", Code: 200}
	json.NewEncoder(w).Encode(hb)
}
