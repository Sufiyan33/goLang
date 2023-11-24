package main

import "fmt"

// Switch statement with an optional statement & expression
func dayPrint() {
	switch day := 4; day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid days")
	}
}

// Switch statement without an optional statement & expression
func print() {
	fmt.Println("-------- Switch statement without an optional statement & expression --------")
	var value int = 2

	switch {
	case value == 1:
		fmt.Println("Hello Go")
	case value == 2:
		fmt.Println("You are amazing")
	case value == 3:
		fmt.Println("Lets explore you")
	default:
		fmt.Println("May be you are looking something else 😒")
	}
}

// Switch statement without default expression
func withoutDefault() {
	var value string = "four"

	switch value {
	case "one":
		fmt.Println("Java")
	case "two":
		fmt.Println("Go")
	case "four":
		fmt.Println("Python")
	}
}
