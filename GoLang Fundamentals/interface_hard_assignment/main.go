package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//fmt.Println(os.Args[1])

	//Open function return a pointer where file has been stored & error
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error is :: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
