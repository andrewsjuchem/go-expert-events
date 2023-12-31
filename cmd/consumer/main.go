package main

import (
	"fmt"

	"github.com/andrewsjuchem/go-expert-events/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)        // Create a Go channel to receive messages from RabbitMQ
	go rabbitmq.Consume(ch, msgs, "orders") // Keeps consuming message from RabbitMQ and puts them into the Go channel

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}
}
