package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloWord(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to GORM")
}

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", helloWord).Methods("GET")
	router.HandleFunc("/user/{name}", NewUser).Methods("POST")
	router.HandleFunc("/users", AllUsers).Methods("GET")
	router.HandleFunc("/user/{name}", UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{delete}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
func main() {
	fmt.Println("Welcome to building Api with GORM")
	initializeDatabse()
	handleRequest()
}
