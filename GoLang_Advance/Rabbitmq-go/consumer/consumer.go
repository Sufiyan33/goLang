package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func throwErrors(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
func main() {
	// First connect to RabbitMQ & then create channel post that create consumer

	fmt.Println("Welcome to goLang with RabbitMQ")
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5671/")
	throwErrors(err, "Failed to connect RabbitMQ")

	defer connection.Close()
	fmt.Println("Successfully connected to RabbitMQ instance")

	// opening a channel over the connection established to interact with RabbitMQ
	channel, err := connection.Channel()
	throwErrors(err, "Failed to open channel")
	defer channel.Close()

	// declaring consumer with its properties over channel opened
	msg, err := channel.Consume(
		"demo", //queue
		"",     //consumer
		true,   //ack
		false,  // exclusive
		false,  // no local
		false,  //no wait
		nil,    // arguments
	)
	throwErrors(err, "Message did not consume")

	// print consumed messages from queue
	myChannel := make(chan bool)
	go func() {
		for message := range msg {
			fmt.Printf("Received message: %s\n", message.Body)
		}
	}()
	fmt.Println("Waiting for messages...")
	<-myChannel
}
