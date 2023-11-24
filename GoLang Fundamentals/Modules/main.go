package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
first create package then create mod file using go mod init.
In order to work with backend and local host, we need a router who will route reques henbce use github.com/gorilla/mux
*/
func main() {
	fmt.Println("Welcome to the section of Modules in goLang")
	greeter()

	// Use route
	router := mux.NewRouter()
	router.HandleFunc("/", serveHome).Methods("GET")

	// Run on server
	log.Fatal(http.ListenAndServe(":4000", router))
}

func greeter() {
	fmt.Println("Hey there mod user 😍")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to goLang section of handle & creating module</h1>"))
}
