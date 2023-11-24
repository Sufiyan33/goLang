package main

func Average(arr [6]int) int {
	var i, value, result int

	for i = 0; i < len(arr); i++ {
		value += arr[i]
	}

	result = value / len(arr)
	return result
}
