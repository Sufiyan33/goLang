package main

import (
	"fmt"
)

/*
The idea is to check how many port is running, if running then return true otherwise false.
*/
func main() {
	fmt.Println("Welcome to scanning port in Go")

	open := port.ScanPort("tcp", "localhost", 1314)
	fmt.Printf("Port open %t\n", open)
}
