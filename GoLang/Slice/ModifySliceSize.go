package main

import "fmt"

// Using append method
func modifySlice() {
	fmt.Println("---------Using append method -------")
	slice := []int{1, 2, 3, 4, 5}
	modifiedSlice := append(slice, 10, 11, 12, 13, 14, 15)

	fmt.Println("Original slice : ", slice)
	fmt.Println("Modify slice : ", modifiedSlice)
}

func modifySliceUsignCopyfunc() {
	fmt.Println("---------Using copy method -------")
	slice := []int{1, 2, 3, 4, 5}
	copySlice := []int{10, 20, 30, 40, 50}
	modifiedSlice := copy(slice, copySlice)

	fmt.Println("Original slice : ", slice)
	fmt.Println("Copy slice : ", copySlice)
	fmt.Println("Modify slice : ", modifiedSlice)
}
