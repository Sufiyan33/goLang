package main

//var deskSize int

func main() {
	//Declaring new variable
	//var card string = "This is a example of variable declarations"

	//You can declare variable like below :
	//panda := "This is panda"
	//Remember, you can't use := if you are re-assigning variable. for this use below :
	//pand = "New panda"

	//deskSize = 52
	//fmt.Println("Hello world!!")
	//fmt.Println(panda)

	//you can call deck class method print here.
	//card.print()

	//Now call deal func here
	//hand, remainingCards := deal(cards, 5)

	//Now by using above deck you can also call print func
	//hand.print()
	//remainingCards.print()

	//Implement []bytes slice
	//greeting := "Hi There!!!"
	//fmt.Println([]byte(greeting))

	// here you can call saveToFile func to save file into local hard drive
	//card.saveToFile("my_cards")

	//You can call here shuffle method to check actually shuffling
	cards := newDeck()
	cards.shuffle()
	deal(cards, 4)
	cards.print()

}
