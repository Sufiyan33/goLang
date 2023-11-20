package main

func swappingByReference(x, y *int) int {
	var temp int

	temp = *x
	*x = *y
	*y = temp
	return temp
}
