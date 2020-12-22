package main

import (
		"encoding/json"
    "log"
    "net/http"
		"github.com/gorilla/mux"
		// "io/ioutil"
		"fmt"
	)

import "github.com/AntonnyAgs/go-api/functions"
import "github.com/AntonnyAgs/go-api/domain"

type CheckConflict struct{
	Agendas domain.Agendas `json:"agendas"`
	Allocations domain.Agendas `json:"allocations"`
}

func HelloWorld(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode("Hello World")
}

func checkConflict(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var checkConflict CheckConflict
	_ = json.NewDecoder(r.Body).Decode(&checkConflict)

	conflicts := functions.CheckConflicts(checkConflict.Agendas, checkConflict.Allocations) 

	fmt.Println(checkConflict.Agendas, checkConflict.Allocations)

	json.NewEncoder(w).Encode(conflicts)
}

func main() {
    router := mux.NewRouter()
		
		router.HandleFunc("/", HelloWorld).Methods("GET")
		router.HandleFunc("/check-conflicts", checkConflict).Methods("POST")

		log.Fatal(http.ListenAndServe(":8000", router))
}
