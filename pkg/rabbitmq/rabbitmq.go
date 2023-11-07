package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel() (*amqp.Channel, error) {
	amqpServerURL := "amqp://guest:guest@localhost:5672/"
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	// defer connectRabbitMQ.Close() // it is closed by the caller

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	// defer channelRabbitMQ.Close() // it is closed by the caller

	// Limits the number of messages that can be processed at the same time
	channelRabbitMQ.Qos(
		100,   // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		panic(err)
	}
	return channelRabbitMQ, nil
}

func Consume(ch *amqp.Channel, out chan<- amqp.Delivery, queueName string) error {
	// Consumes messages from RabbitMQ
	msg, err := ch.Consume(
		queueName,     // queue
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

func Publish(ch *amqp.Channel, body string, exName string, queueName string) error {
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	}
	err := ch.Publish(
		exName,
		queueName,
		false,
		false,
		message,
	)
	if err != nil {
		return err
	}
	return nil
}
