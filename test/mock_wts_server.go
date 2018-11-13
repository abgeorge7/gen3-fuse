package main

import (
	"fmt"
	"log"
	//"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func auth_url_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

	params := mux.Vars(r)
	authorization_URL := params["authorization_url"]

	fmt.Println(authorization_URL)

}

func token_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	// fmt.Println(r.Form)

	expires, ok := r.URL.Query()["expires"]

	if !ok || len(expires[0]) < 1 {
		log.Println("Url Param 'expires' is missing")
		return
	}

	num_seconds := expires[0]

	fmt.Println(num_seconds)

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/oauth2/{authorization_url}", auth_url_handler).Methods("GET")
	router.HandleFunc("/token", token_handler).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:8001", router))
}
