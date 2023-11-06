package main

import (
	"fmt"
)

type EventInterface interface {
	GetName() string
}

type Event struct {
	Name string
}

func (e *Event) GetName() string {
	return e.Name
}

func Dispatch(event EventInterface) error {
	fmt.Println(event.GetName())
	return nil
}

func main() {
	var createOrderEvent = Event{Name: "Create order"}
	fmt.Println("*** FIRST ORDER ***")
	/*
		The createOrderEvent variable is passed as a pointer to the Dispatch function because of the way the EventInterface is implemented.
		When defining an interface in Go, the methods that make up the interface can be implemented using either value receivers or pointer receivers.
		In this instance, the GetName method of the Event struct is implemented with a pointer receiver (e *Event),
		which means that only pointers to Event instances can be used as EventInterface.
	*/
	Dispatch(&createOrderEvent)
}
