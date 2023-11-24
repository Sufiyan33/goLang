package main

import "fmt"

var url string = "www.google.com"

const FIRST_NAME = "Creating Constant"

func main() {
	var username = "sufiyan"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var value int = 245
	fmt.Println(value)
	fmt.Printf("Variable is of type: %T \n", value)

	var isTrue bool = false
	fmt.Println(isTrue)
	fmt.Printf("Variable is of type: %T \n", isTrue)

	// Using walrus operator
	name := "vs code"
	fmt.Println(name)
	fmt.Printf("Variable is of type: %T \n", name)

	floatValue := 2947914.247279
	fmt.Println(floatValue)
	fmt.Printf("Variable is of type: %T \n", floatValue)

	// Accessing gloabl variables
	fmt.Println(url)
	fmt.Printf("Variable is of type: %T \n", url)

	fmt.Println(FIRST_NAME)
	fmt.Printf("Variable is of type: %T \n", FIRST_NAME)

}
