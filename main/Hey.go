package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)


func test(w http.ResponseWriter, r *http.Request)  {



	json.NewEncoder(w).Encode(struct { // anonymous struct
		ID string
	}{ID: "555"})
}

func main() {

	fmt.Print("hel")
	router := mux.NewRouter()

	router.HandleFunc("/test",test)

	http.ListenAndServe(":5000",router)
	
}
