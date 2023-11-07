package main

import (
	"fmt"

	"github.com/andrewsjuchem/go-expert-events/pkg/rabbitmq"
)

func main() {
	// Create RabbitMQ channel (this is not a regular Go channel)
	channelRabbitMQ, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	_, err = channelRabbitMQ.QueueDeclare(
		"orders", // queue name
		true,     // durable
		false,    // auto delete
		false,    // exclusive
		false,    // no wait
		nil,      // arguments
	)
	if err != nil {
		panic(err)
	}

	// Produce multiple messages to RabbitMQ
	for i := 0; i < 50; i++ {
		var message = fmt.Sprintf("Hello World %d!", i+1)
		err = rabbitmq.Publish(channelRabbitMQ, message, "", "orders")
		if err != nil {
			panic(err)
		}
		fmt.Println(message)
	}
}
