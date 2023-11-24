package main

import "fmt"

func checkSliceEmpty() {
	r := []int{1, 2, 3}
	if len(r) == 0 {
		fmt.Println("Slice is empty !")
	} else {
		fmt.Println("Slice is not empty 😊")
	}
}
