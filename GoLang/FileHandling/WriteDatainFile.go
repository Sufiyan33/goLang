package main

import (
	"fmt"
	"io"
	"os"
)

func writeData() {
	content := "Welcode to the File handling in Go... let see how to deal with this"

	file, err := os.Create("./demo.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("Length of the content written in file are :", length)
	defer file.Close()
}
