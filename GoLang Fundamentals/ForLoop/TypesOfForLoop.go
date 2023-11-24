package main

import "fmt"

var weekDays = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

func normalForLoop() {
	fmt.Println("----- for loop as a normal for loop -----")
	for d := 0; d < len(weekDays); d++ {
		fmt.Println(weekDays[d])
	}
}

func forLoopAsWhile() {
	fmt.Println("----- for loop as a while loop -----")
	var index = 1
	for index < len(weekDays) {
		fmt.Println(weekDays[index])
		index++
	}
}

func forLoopAsRange() {
	fmt.Println("----- for loop as a range -----")
	for index, days := range weekDays {
		fmt.Printf("Index is %v & days is %v \n", index, days)
	}
}

func forLoopForStrings() {
	fmt.Println("----- for loop for strings -----")

	for index, str := range "Sufiyan" {
		fmt.Printf("Index is %v & unicode of string is %v\n", index, str)
	}
}

func forLoopforMap() {
	fmt.Println("----- for loop for map -----")

	mapData := map[int]string{
		1: "Go",
		2: "Java",
		3: "C++",
		4: "C",
		5: "Pyhton",
		6: "JavaScript",
		7: "Ruby",
	}
	for key, value := range mapData {
		fmt.Printf("Key is %v, & value of map is %v\n", key, value)
	}
}

func forLoopForChannel() {
	fmt.Println("----- for loop for channel -----")

	myChannel := make(chan int)
	go func() {
		myChannel <- 10
		myChannel <- 100
		myChannel <- 1000
		myChannel <- 10000
		myChannel <- 100000
		close(myChannel)
	}()
	for i := range myChannel {
		fmt.Println("Value that channel received are: \n", i)
	}
}
