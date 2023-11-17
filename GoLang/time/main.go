package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome to time study of goLang 🙂")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006"))
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	fmt.Println("----- Lets create custome data -----")
	customeDate := time.Date(2020, time.January, 12, 10, 23, 0, 0, time.UTC)
	fmt.Println(customeDate)
	fmt.Println(customeDate.Format("01-02-2006 Monday "))
}
