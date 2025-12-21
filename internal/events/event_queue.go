package events

import (
	"log"
	"sync"
)

type EventQueue struct {
	eventChan chan Event
	mutex     sync.Mutex
}

func NewEventQueue() *EventQueue {
	return &EventQueue{
		eventChan: make(chan Event),
		mutex:     sync.Mutex{},
	}
}

func (q *EventQueue) Enqueue(event Event) {
	log.Printf("enqueue event %s", event.Type())
	q.eventChan <- event
}

func (q *EventQueue) Dequeue() Event {
	return <-q.eventChan
}

func (q *EventQueue) Close() {
	close(q.eventChan)
}

func (q *EventQueue) Length() int {
	return len(q.eventChan)
}

func (q *EventQueue) Loop() {
	for {
		event := q.Dequeue()
		err := event.Execute()
		if err != nil {
			panic(err)
		}
	}
}
