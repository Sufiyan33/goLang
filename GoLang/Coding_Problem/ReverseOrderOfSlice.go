package main

func reverse(sw []int) {
	for a, b := 0, len(sw)-1; a < b; a, b = a+1, b-1 {
		sw[a], sw[b] = sw[b], sw[a]
	}
}
