package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

// Defining variable of type websocket
var upgrader = websocket.Upgrade{}
var todoList []string

func getCmd(input string) string {
	inputArr := strings.Split(input, " ")
	return inputArr[0]
}

func getMessage(input string) string {
	inputArr := strings.Split(input, " ")
	var result string
	for i := 1; i < len(inputArr); i++ {
		result += inputArr[i]
	}
	return result
}

func updateTodoList(input string) {
	tempList := todoList
	todoList = []string{}
	for _, val := range tempList {
		if val == input {
			continue
		}
		todoList = append(todoList, val)
	}

}

func main() {
	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Setting up the server 😍")
		// Upgrade upgrades the HTTP server connection to the WebSocket protocol.

		con, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade failed 🤯", err)
			return
		}
		defer con.Close()

		// Continuosly read and write message
		for {
			mt, message, err := con.ReadMessage()
			if err != nil {
				log.Print("read failed", err)
				break
			}
			input := string(message)
			cmd := getCmd(input)
			msg := getMessage(input)

			if cmd == "add" {
				todoList = append(todoList, cmd)
			} else if cmd == "done" {
				updateTodoList(msg)
			}
			output := "Current Todos; \n"
			for _, todo := range todoList {
				output += "\n - " + todo + "\n"
			}

			output += "\n-------------------------"
			message = []byte(output)
			err = con.WriteMessage(mt, message)
			if err != nil {
				log.Print("Write failed", err)
				break
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "WebSockets.html")
	})
	http.ListenAndServe(":8000", nil)
}
