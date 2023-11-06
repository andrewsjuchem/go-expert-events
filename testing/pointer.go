package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// ---------------
type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return errors.New("handler already registered")
			}
		}
	}
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}

func (ev *EventDispatcher) Dispatch(event EventInterface) error {
	if handlers, ok := ev.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			go handler.Handle(event, wg)
		}
		wg.Wait()
	}
	return nil
}

// ---------------
type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
}

type Event struct {
	Name string
}

func (e *Event) GetName() string {
	return e.Name
}

func (e *Event) GetDateTime() time.Time {
	return time.Now()
}

// ---------------
type EventHandler struct {
	message string
}

func (h *EventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	fmt.Println(h.message)
	wg.Done()
}

// ---------------

func main() {
	var eventDispatcher = NewEventDispatcher()
	var applyDiscountHandler = EventHandler{message: "A discount has been applied."}
	var createOrderEvent = Event{Name: "Create order"}

	err := eventDispatcher.Register(createOrderEvent.GetName(), &applyDiscountHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("*** FIRST ORDER ***")
	eventDispatcher.Dispatch(&createOrderEvent)
}
