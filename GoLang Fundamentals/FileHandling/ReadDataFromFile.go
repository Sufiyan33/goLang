package main

import (
	"fmt"
	"io/ioutil"
)

func readData() {
	fmt.Println("Reading data from file")
	dataType, err := ioutil.ReadFile("./demo.txt")
	checkNilError(err)

	//Whenever you will read file, it will give you in byte form. You need to convert it.
	fmt.Println("Reading content of file & that are :", dataType)
	fmt.Println("Reading content of file & that are :", string(dataType))
}
