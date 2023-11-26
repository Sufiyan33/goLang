package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	fmt.Println("Hello websocket")

	server := socketio.NewServer(nil)

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)

		s.join("chat")

	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.Handle("/socket.io/", server)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
