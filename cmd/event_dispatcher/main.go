package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/andrewsjuchem/go-expert-events/pkg/events"
)

type Event struct {
	Name    string
	Payload interface{}
}

type EventPayload struct {
	ID       int
	Price    float64
	Discount float64
}

func (e *Event) GetName() string {
	return e.Name
}

func (e *Event) GetPayload() interface{} {
	return e.Payload
}

func (e *Event) GetDateTime() time.Time {
	return time.Now()
}

type EventHandler struct {
	message string
}

func (h *EventHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	fmt.Println(h.message)
	wg.Done()
}

func main() {
	var eventDispatcher = events.NewEventDispatcher()
	var applyDiscountHandler = EventHandler{message: "A discount has been applied."}
	var calculateOrderHandler = EventHandler{message: "Order has been calculated."}
	var confirmPaymentHandler = EventHandler{message: "Payment has been confirmed."}

	var createOrderEvent = Event{Name: "Create order", Payload: EventPayload{ID: 1, Price: 60, Discount: 10}}

	err := eventDispatcher.Register(createOrderEvent.GetName(), &applyDiscountHandler)
	if err != nil {
		panic(err)
	}

	err = eventDispatcher.Register(createOrderEvent.GetName(), &calculateOrderHandler)
	if err != nil {
		panic(err)
	}

	err = eventDispatcher.Register(createOrderEvent.GetName(), &confirmPaymentHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("*** FIRST ORDER ***")
	eventDispatcher.Dispatch(&createOrderEvent)

	fmt.Println("*** SECOND ORDER ***")
	err = eventDispatcher.Remove(createOrderEvent.GetName(), &confirmPaymentHandler)
	if err != nil {
		panic(err)
	}
	eventDispatcher.Dispatch(&createOrderEvent)
}
