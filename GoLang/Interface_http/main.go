package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//Creating custom writer fucntion
type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error is :: ", err)
		os.Exit(1)
	}

	//making byte slice using in built in function
	/*bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))*/

	//you can also write code like below :
	//io.Copy(os.Stdout, resp.Body)

	//you can call Write method here
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
