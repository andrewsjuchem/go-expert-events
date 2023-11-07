package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel() (*amqp.Channel, error) {
	// Create connection to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		panic(err)
	}
	// defer conn.Close()

	// Create RabbitMQ channel (this is not a regular Go channel)
	ch, err := conn.Channel()
	// Limits the number of messages that can be processed at the same time
	ch.Qos(
		100,   // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		panic(err)
	}
	// defer ch.Close()
	return ch, nil

}

func Consume(ch *amqp.Channel, out chan<- amqp.Delivery, queue string) error {
	// Consumes messages from RabbitMQ
	msg, err := ch.Consume(
		queue,         // queue
		"go-consumer", // consumer
		false,         // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	if err != nil {
		return err
	}
	// Reads every message that is being consumed and puts into the Go channel
	for msg := range msg {
		out <- msg
	}
	return nil
}

func Publish(ch *amqp.Channel, body string, exName string) error {
	ctx := context.Background()
	err := ch.PublishWithContext(
		ctx,
		exName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		return err
	}
	return nil
}
