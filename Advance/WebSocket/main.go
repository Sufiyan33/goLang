package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Setting up the server 😍")
	})
	http.ListenAndServe(":8000", nil)
	//con, err := upgrader.Upgrade(w, r, nil)
}
