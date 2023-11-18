package main

import (
	"fmt"
	"math/rand"
	"time"
)

func diceGame() {
	fmt.Println("Welcome to the Game of Dice")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 & you can open")
	case 2:
		fmt.Println("You can move to 2 spot")
	case 3:
		fmt.Println("You can move to 3 spot")
	case 4:
		fmt.Println("You can move to 4 spot")
	case 5:
		fmt.Println("You can move to 5 spot")
	case 6:
		fmt.Println("You can move 6 spot and roll dice again")
	default:
		fmt.Println("Sorry wrong number")
	}
}
