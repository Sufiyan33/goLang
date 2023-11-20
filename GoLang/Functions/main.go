package main

import "fmt"

func main() {
	fmt.Println("Area is :: ", area(12, 12))

	fmt.Println("----- Call by value -----")
	var a int = 10
	var b int = 20
	fmt.Printf("a = %d and b = %d", a, b)
	_ = swap(a, b)
	fmt.Printf("\na = %d, and b = %d", a, b)

	fmt.Println("\n----- Call by reference -----")
	var x int = 30
	var y int = 40
	fmt.Printf("x = %d and y = %d", x, y)
	swappingByReference(&x, &y)
	fmt.Printf("\nx = %d and y = %d", x, y)

	fmt.Println("\n----- Variadic Function -----")
	// Passing zero arguments to function
	fmt.Println(joinStrings())

	// Passing multiple parameters to the function
	fmt.Println(joinStrings("Sufiyan", "Ahmad"))
	fmt.Println(joinStrings("Go", "is"))
	fmt.Println(joinStrings("the", "best"))
	fmt.Println(joinStrings("language", "to", "learn"))

	// Passing slice as parameters
	stringSlice := []string{"Welcome", "to", "the", "world", "of", "go", "language"}
	fmt.Println(joinStrings(stringSlice...))

	// Anonymous Fucntion
	fmt.Println("----- Anonymous Function -----")
	anonymousExample1()
	anonymousExample2()
	anonymousExample3()
	value := func(x, y string) string {
		return x + y + "GoLanguage"
	}
	anonymousExample4(value)

	result := anonymousExample5()
	fmt.Println(result("Welcome ", "to "))
}
