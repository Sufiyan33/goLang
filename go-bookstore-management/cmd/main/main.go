package main

import (
	"log"
	"net/http"

	"github.com/Sufiyan33/goLang/go-bookstore-management/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	homePage := http.FileServer(http.Dir("./cmd"))
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	http.Handle("/home", homePage)
	log.Fatal(http.ListenAndServe("localhost:9000", r))
}
