package main

import "github.com/andrewsjuchem/go-expert-events/pkg/rabbitmq"

func main() {
	// Create RabbitMQ channel (this is not a regular Go channel)
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "Hello World!", "amq.direct")

	// for i := 0; i < 10000000; i++ {
	// 	order := GenerateOrders()
	// 	err := Notify(ch, order)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(order)
	// }
}
