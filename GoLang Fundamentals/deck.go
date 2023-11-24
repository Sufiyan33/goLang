package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new type of 'deck'
//which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Club"}
	cardValues := []string{"One", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" "+value)
		}
	}
	return cards
}

//create a function and call this in main class
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Create a function and return multiple values
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//Converting deck to []byte
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//Save this file into local hard drive
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

//Read file from local hard drive
func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	//Option 1 : Log the error & return a call to newDeck
	//Option 2 : log the error & entirely quite the program
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1) //means program is not success and quite program
	}

	//Here, Line 51 returning bs (means byte slices), we need to convert it into strings then deck
	s := strings.Split(string(bs), ",") //this is a string, now convert in deck
	return deck(s)
}

//Create func to shuffle cards
func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		//Now do swapping
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
