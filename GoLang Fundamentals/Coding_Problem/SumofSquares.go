package main

import (
	"fmt"
)

//Question :
/*
Implement the SumOfSquares function which takes an integer, c and returns the sum of all squares between 1 and c.
	You’ll need to use select statements, goroutines, and channels.
For example, entering 5 would return 55 because :
	1^2 + 2^ + 3^ + 4^ + 5^ = 55
*/

func doSum() {
	myChannel := make(chan int)
	quiteChannel := make(chan int)

	sum := 0
	go func() {
		for i := 1; i <= 5; i++ {
			sum += <-myChannel
		}
		fmt.Println(sum)
		quiteChannel <- 0
	}()
	sumOfSquare(myChannel, quiteChannel)
}

func sumOfSquare(c, quite chan int) {
	y := 1
	for {
		select {
		case c <- (y * y):
			y++
		case <-quite:
			return
		}
	}
}
