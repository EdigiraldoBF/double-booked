package eventsloader

import "github.com/EdigiraldoBF/double-booked/events"

type EventsLoader interface {
	Load() (events.Events, error)
}

var eventsLoader EventsLoader

func SetLoader(concreteLoader EventsLoader) {
	eventsLoader = concreteLoader
}

func Load() (events.Events, error) {
	return eventsLoader.Load()
}
