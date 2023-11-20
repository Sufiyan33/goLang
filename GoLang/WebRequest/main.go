package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "http://www.google.com"

func main() {
	fmt.Println("Welcome to handling of web request in Go")

	response, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	dataTypes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dataTypes))
}
