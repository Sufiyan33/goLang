package main

import "fmt"

func main() {
	// Creating 1-D array
	arr := [3]string{"Sufiyan", "Ahmad", "Go"} //Using shorthand
	fmt.Println(arr)

	var arr1 [3]string //Using var keyword
	arr1[0] = "Go"
	arr1[1] = "is"
	arr1[2] = "amazing"
	fmt.Println("Using var keyword", arr1)
	for x := 0; x < len(arr1); x++ {
		fmt.Println(arr1)
	}

	// You can use for loop to access array elements.
	fmt.Println("Elements of arrays are :: ")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// Creating 2-D array
	two := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("2-D Elements of the array are :: ")
	fmt.Println(two)

	for i := 0; i < len(two); i++ {
		for j := 0; j < len(two); j++ {
			fmt.Println(two[i][j])
		}
	}

	// You can also use ellipsis instead of length

	name := [...]string{"sufiyan", "ahmad,", "go", "is", "amazing"}
	fmt.Println("Elements of array are :: ", name)
	fmt.Println("Size of the array are :: ", len(name))

	// Array is of type value not of type reference.
	array := [...]int{1, 2, 3, 4, 12, 14, 15}
	fmt.Println("Original array :: ", array)

	// assigning to new variable
	myArray := array
	fmt.Println("new array :: ", myArray)
	// Making changes
	myArray[0] = 50
	fmt.Println("Made changes in new array :: ", myArray)
	fmt.Println("Original array :: ", array)

	fmt.Println("--------- Copy an arrya into another array ---------")
	fmt.Println("Approach 1 :: By Passing value")

	my_array1 := [3]string{"Soabe", "Akhtar", "Sabir"}
	//Assiging this into another array by using value
	my_array2 := my_array1

	fmt.Println("my_Array1 :: ", my_array1)
	fmt.Println("my_Array2 :: ", my_array2)

	// Now making some changes.
	my_array1[0] = "Sufiyan"

	// Here copying into another and making changes in original array. It will not reflect in copy(my_array2) array.
	fmt.Println("my_Array1 :: ", my_array1)
	fmt.Println("my_Array2 :: ", my_array2)

	fmt.Println("Approach 2 :: By Passing reference")

	my_array3 := [4]int{12, 13, 14, 15}

	// Copying array into another array by reference
	my_array4 := &my_array3

	fmt.Println("my_array 3 :: ", my_array3)
	fmt.Println("my_array 4 :: ", *my_array4)

	// Now making some changes in original array
	my_array3[3] = 16

	fmt.Println("my_array 3 :: ", my_array3)
	fmt.Println("my_array 4 :: ", *my_array4)

	fmt.Println("--------- Multiple ways to Copy an arrya into another array ---------")
	fmt.Println("Approach 1 :: By using a loop ")

	originalArray := [...]int{1, 2, 3, 4, 5, 6}
	copyArray := make([]int, len(originalArray))

	for i, value := range originalArray {
		copyArray[i] = value
	}
	fmt.Println("Original array :: ", originalArray)
	fmt.Println("Copy array :: ", copyArray)

	fmt.Println("Approach 2 :: By using copy function ")

	originalArray2 := []int{1, 2, 3, 4, 5, 6}
	copyArray2 := make([]int, len(originalArray))

	copy(copyArray2, originalArray2)

	fmt.Println("Original array :: ", originalArray2)
	fmt.Println("Copy array :: ", copyArray2)

	fmt.Println("Approach 3 :: By using append function ")

	originalArray3 := []int{1, 2, 3, 4, 5, 6}
	copyArray3 := make([]int, 0, len(originalArray))

	copyArray3 = append(copyArray3, originalArray3...)

	fmt.Println("Original array :: ", originalArray3)
	fmt.Println("Copy array :: ", copyArray3)

	fmt.Println("------- Find Average of an array --------")
	ar := [6]int{10, 20, 30, 40, 50, 60}
	var result = Average(ar)
	fmt.Println("Average of an array is :: ", result)

	fmt.Println("------ Passing an array to function -------")
	aa := [5]int{2, 3, 4, 5, 6}
	printArray(aa)
}

func printArray(aa [5]int) {
	for _, value := range aa {
		fmt.Println(value)
	}
}
