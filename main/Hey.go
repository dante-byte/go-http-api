package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)






var data []string = []string{} // composite literal string slice

func main() {


	router := mux.NewRouter()

	router.HandleFunc("/test",test)


	router.HandleFunc("/add/{item}", addItem)

	http.ListenAndServe(":5000",router)
	
}

// route handlers

func test(w http.ResponseWriter, r *http.Request)  {

	w.Header().Set("Content-Type", "application/json")



	json.NewEncoder(w).Encode(struct { // anonymous struct, a predefined type isn't needed
		ID string //fields
	}{ID: "5"}) // value set
}

func addItem(w http.ResponseWriter, r *http.Request)  {

	routeVar := mux.Vars(r)["item"] // gives access to data
	data = append(data, routeVar)

	json.NewEncoder(w).Encode(data)


}
