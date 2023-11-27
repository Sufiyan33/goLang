package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Welcome to jwt with go")
	http.HandleFunc("/login", Login)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/refresh", Refresh)
	http.HandleFunc("/logout", Logout)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
