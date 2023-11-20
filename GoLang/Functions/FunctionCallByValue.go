package main

func swap(a, b int) int {
	var temp int

	temp = a
	a = b
	b = temp
	return temp
}
