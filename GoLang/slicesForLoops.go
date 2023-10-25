package main

import "fmt"

func slicesForLoops() {
	cards := []string{"Slices-1", "slcies-2", "Slices-4"}
	cards = append(cards, "slices-3")

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

/*func newCard() string {
	return "exmple of slices & for loops"
}*/
