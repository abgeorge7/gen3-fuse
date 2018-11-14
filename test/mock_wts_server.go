package main

import (
	//"encoding/json"
	//"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func auth_url_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{'success': 'connected with fence'}"))
}

func token_handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

	// expires, ok := r.URL.Query()["expires"]

	// if !ok || len(expires[0]) < 1 {
	// 	log.Println("Url Param 'expires' is missing")
	// 	return
	// }

	// // num_seconds := expires[0]

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{'token': 'MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI3'}"))

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/oauth2/authorization_url", auth_url_handler).Methods("GET")
	router.HandleFunc("/token", token_handler).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:8001", router))
}
