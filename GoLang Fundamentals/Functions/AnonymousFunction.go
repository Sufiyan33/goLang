package main

import "fmt"

func anonymousExample1() {
	func() {
		fmt.Println("Welcome to Anonymous function")
	}()
}

func anonymousExample2() {
	value := func() {
		fmt.Println("Assigning an anonymous function to a variable.")
	}
	value()
}

func anonymousExample3() {
	func(str string) {
		fmt.Println("Passing Arguments to an Anonymous function")
		fmt.Println(str)
	}("Passing arguments")
}

func anonymousExample4(i func(x, y string) string) {
	fmt.Println("Passing an Anonymous function as an argument to other function")
	fmt.Println(i("go", "lang"))
}

func anonymousExample5() func(i, j string) string {
	fmt.Println("Returning an anonymous function from another function")

	myFunc := func(i, j string) string {
		return i + j + "Returning anonymous function"
	}
	return myFunc
}
