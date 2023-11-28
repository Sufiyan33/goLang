package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func throwError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
func main() {
	fmt.Println("Welcome to goLang with RabbitMQ")
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5671/")
	throwError(err, "Failed to connect RabbitMQ")

	defer connection.Close()
	fmt.Println("Successfully connected to RabbitMQ instance")

	// opening a channel over the connection established to interact with RabbitMQ
	channel, err := connection.Channel()
	throwError(err, "Failed to open channel")
	defer channel.Close()

	// declaring queue with its properties over the the channel opened
	queue, err := channel.QueueDeclare(
		"demo", // queue name
		false,  //durable
		false,  //delete when not used
		false,  //exclusive
		false,  //no-wait
		nil,    //arguments
	)
	throwError(err, "Failed to declare Queue")

	// publishing a message
	err = channel.Publish(
		"",        // exchange
		"testing", //key
		false,     //mandatory
		false,     //immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello, from RabbitMQ"),
		},
	)
	throwError(err, "Message is not publish")

	fmt.Println("Quesue status :: ", queue)
	fmt.Println("Message succesfully published")
}
