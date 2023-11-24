package main

import "fmt"

func sliceFromArray() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:4]

	fmt.Println("Array : ", array)
	fmt.Println("Slice : ", slice)
}
