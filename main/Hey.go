package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)




type Item struct {

	data string
}


var data []Item = []Item{} // composite literal string slice

func main() {


	router := mux.NewRouter()

	router.HandleFunc("/add", addItem).Methods("POST") //{} dynamive varial identifer value .METHODS ALLOWS CERTER HTTP REUEST
	http.ListenAndServe(":5000",router)
	
}

// route handlers


func addItem(w http.ResponseWriter, r *http.Request)  { // * pointer type https.reqyest
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data) // encode the data from teh slice

	var newItem Item
	json.NewDecoder(r.Body).Decode(&newItem)


	data = append(data, newItem)


}
