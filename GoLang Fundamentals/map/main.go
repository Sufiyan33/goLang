package main

import "fmt"

func main() {

	//way 1 : to declare map
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#78ff1ba",
		"white": "#ft0142ba",
	}
	//fmt.Println(colors)
	//call PrintMap function here
	printMap(colors)

	//way 2 : to declare map
	var ball map[string]string
	fmt.Println(ball)

	//way 3 : to declare map
	home := make(map[string]string)
	fmt.Println(home)

	//Manipulating
	home["Mannat"] = "House"
	home["Galaxy"] = "Apartment"
	home["Ajay"] = "Complex"
	fmt.Println(home)

	//delete operation
	delete(home, "Mannat")

	//Question : related to change in map
	m := map[string]string{
		"dog": "bark",
	}
	changeMap(m)
	fmt.Println(m)

	//Question : can we iterate a map without a function.
	//Answer : yes, can do like below.
	animal := map[string]string{
		"dog":  "bark",
		"cat":  "purr",
		"goat": "may",
		"lion": "roar",
	}

	for key, value := range animal {
		fmt.Println(key, value)
	}
}

//Writing a function to iterate map
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, "is ", hex)
	}
}

func changeMap(m map[string]string) {
	m["cat"] = "purr"
}
