package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting API...")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", heartbeat)
	log.Fatal(http.ListenAndServe(":8080", router))

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
