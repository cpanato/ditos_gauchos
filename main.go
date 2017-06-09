package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cpanato/ditos_gauchos/ditos"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/bah", handleBah).Methods("POST")

	ditos.LoadDitos()

	fmt.Println("Running")

	http.ListenAndServe(":8080", router)

}

func handleBah(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	bahMessage, _ := ditos.GenerateRandomDito()

	outgoingJSON, error := json.Marshal(bahMessage)

	if error != nil {
		log.Println(error.Error())
		http.Error(res, error.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(res, string(outgoingJSON))
}
