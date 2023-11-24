package main

import "fmt"

type Address struct {
	name    string
	city    string
	state   string
	pinCode int
}

// Creating struct with anonymous fields
type Student struct {
	int
	string
	float64
}

func main() {
	//declaring a variable of struct type. So all variable will be assign to their default values.
	var a Address
	fmt.Println("Values at the time of declaring structs : ", a)

	// Declaring & initializing a struct using struct literal
	a1 := Address{"Sufiyan", "Bangalore", "Karnatka", 00000}
	fmt.Println("Address 1 : ", a1)

	//Accessing struct variable by using dot.
	fmt.Println("Name : ", a1.name)
	fmt.Println("City : ", a1.city)
	fmt.Println("State : ", a1.state)
	fmt.Println("Pin code : ", a1.pinCode)

	// Naming fields while initializing struct.
	a2 := Address{name: "Sufiyan", city: "Bangalore", state: "karnataka", pinCode: 00000}
	fmt.Println("Address 2 : ", a2)

	//Not assigning a value
	a3 := Address{name: "Sufiyan"}
	fmt.Println("Address 3 : ", a3)

	// Creating and initializing anonymous structure.
	element := struct {
		name      string
		branch    string
		language  string
		particles int
	}{
		name:      "sufiyan",
		branch:    "CSE",
		language:  "Go",
		particles: 475,
	}
	fmt.Println(element)

	// Assigning value to anonymous struct with anonymous fields.
	value := Student{123, "Sufiyan", 80.08}
	fmt.Println("Anonymous struct : ", value)

	fmt.Println("Student Id : ", value.int)
	fmt.Println("Student Name : ", value.string)
	fmt.Println("Book Price : ", value.float64)
}
