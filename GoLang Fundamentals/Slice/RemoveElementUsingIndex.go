package main

import "fmt"

func removeElementBasedonIndex() {

	elements := []string{"Java", "C", "Python", "Go", "Swift", "Android"}
	fmt.Println("original slice: ", elements)

	var index int = 2
	elements = append(elements[:index], elements[index+1:]...)
	fmt.Println("Slice after removed one index: ", elements)
}
