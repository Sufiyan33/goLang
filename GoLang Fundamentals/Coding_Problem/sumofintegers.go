package main

import "fmt"

func sumofNumber() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum of number from 1 to 10", sum)
}

func sumofTwoNumber(x, y int) int {
	sum := 0
	sum = x + y
	return sum
}

func addNumbers() {
	number := []int{1, 3, 4, 5, 6, 7, 10}

	for _, num := range number {
		num += num
		fmt.Println(num)
	}
}
