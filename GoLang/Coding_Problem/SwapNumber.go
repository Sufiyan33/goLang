package main

func swap() []int {
	a, b := 10, 15
	b, a = a, b

	return []int{a, b}
}
