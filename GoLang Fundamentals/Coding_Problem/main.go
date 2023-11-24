package main

import "fmt"

func main() {
	sumofNumber()

	numberSum := sumofTwoNumber(10, 20)
	fmt.Println("Two number sum is : ", numberSum)

	addNumbers()
	implement_stack()

	//swapping number
	fmt.Println("swapped number are : ", swap())

	//Call min & max function
	fmt.Println("min number is :: ", min(5, 10))
	fmt.Println("max number is :: ", max(5, 10))

	// Reverse slice
	x := []int{1, 2, 3}
	reverse(x)
	fmt.Println("Reverse order of slice is :: ", x)

	// check Slice is empty or not
	checkSliceEmpty()

	//merge sort
	data := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}
	fmt.Printf("%v\n%v\n", data, mergeSort(data))
	fmt.Printf("%v\n%v\n", data, mergeSort(data))

	//merge sort using go-routing & channel
	fmt.Printf("%v\n%v\n", data, mergeSortUsingChannel(data))

	//Sum of squares
	doSum()
}
