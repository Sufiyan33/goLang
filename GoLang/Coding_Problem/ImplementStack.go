package main

import "fmt"

func implement_stack() {

	// Create stack using slices
	var stack []string

	// push element

	stack = append(stack, "How are you bro !!!")
	stack = append(stack, "I am fine, what about you?")

	for len(stack) > 0 {
		// print top means peek elements
		n := len(stack) - 1
		fmt.Println(stack[n])

		//pop element
		stack = stack[:n]
	}
}
