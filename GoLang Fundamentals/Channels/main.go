package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://midium.com",
	}

	for _, link := range links {
		//Added sleep() method so that it do not block other responses.
		//time.Sleep(1 * time.Second)
		// commenting sleep() method & creating channel to do communication.
		c := make(chan string)
		go checkLink(link, c)
	}
	c := make(chan string)
	//Creating Routine
	//go checkLink(link, c)
	//fmt.Println(<-c)

	//Let's for loop here to print channel value
	/*for i := 0; i < len(link); i++ {
		go checkLink(<-c, c)
	}*/

	//We can also write for like below  :
	/*for {
		go checkLink(<-c, c)
	}*/

	//We can also write for like below  :
	for l := range c {
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " might be down!!!")
		//c <- "Might be down I think but not sure!"

		//You can also add above code like below :
		c <- link
		return
	}

	fmt.Println(link + " is up😊")
	//c <- "Yup it's up"

	//You can also add above code like below :
	c <- link
}
