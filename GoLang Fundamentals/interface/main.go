package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

/*Here, I am trying producing issue if you are not using interface.
Clearly visible that printGreeting function not allowed to declared multiple time having different types of argument.
While other languages provide overriding or overLaoding concept but here not. Let's see how to deal with this*/

/*Add interface here typ bot and problem will go*/
func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	/*-------------------------*/
	//Accessing element of tank interface
	//var t tank
	//t = myValue{10, 14}
	//fmt.Println("Area of tank :: ", t.Tarea())
	//fmt.Println("Volume of tank :: ", t.Volume())
}

func (englishBot) getGreeting() string {
	return "Hi! there😊"
}

func (spanishBot) getGreeting() string {
	return "Hello"
}

/*func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}*/

/*func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}*/

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
