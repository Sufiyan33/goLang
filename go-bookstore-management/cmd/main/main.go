package main

import (
	"log"
	"net/http"

	"github.com/Sufiyan33/goLang/go-bookstore-management/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRouters(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("9000", r))
}
