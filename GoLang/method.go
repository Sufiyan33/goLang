package main

import "fmt"

func method() {

	// calling function new card here and storing value in card variable
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Declaring method here"
}
