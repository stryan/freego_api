package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	api := NewAPI()
	router := mux.NewRouter().StrictSlash(true)
	//plot out routes
	router.HandleFunc("/game", api.NewGame).Methods("POST")
	router.HandleFunc("/game/{id}", api.GetGame).Methods("GET")
	router.HandleFunc("/game/{id}/status", api.GetGameStatus).Methods("GET")
	router.HandleFunc("/game/{id}/move", nil).Methods("POST")
	router.HandleFunc("/game/{id}/move/{movenum}", nil).Methods("GET")
	log.Fatal(http.ListenAndServe(":1379", router))

}
