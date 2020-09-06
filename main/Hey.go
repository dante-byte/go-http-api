package main

import (
	"fmt"
	"github.com/gorilla/mux"
	//"encoding/json"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request)  {

	w.Write([]byte("this is a test"))

}

func main() {

	fmt.Print("hel")
	router := mux.NewRouter()

	router.HandleFunc("/test",test)

	http.ListenAndServe(":5000",router)
	
}
