package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Description"`
	Content string `json:"Content"`
}
type Articles []Article

func AllArticales(w http.ResponseWriter, r *http.Request) {
	article := Articles{
		Article{
			Title:   "REST-API",
			Desc:    "Learning rest api",
			Content: "You can also use GORM and Gorilla to do this",
		},
	}
	fmt.Println("All articles are")
	json.NewEncoder(w).Encode(article)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page endpoints")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/article", AllArticales)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
func main() {
	handleRequest()
}
