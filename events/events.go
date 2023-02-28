package events

import (
	"time"
)

type Events struct {
	Events []Event
}

type Event struct {
	Start time.Time
	End   time.Time
}

func (e Event) OverlapsWith(anotherEvent Event) bool {
	if e.Start.Equal(anotherEvent.Start) && e.End.Equal(anotherEvent.End) {
		return true
	}

	if e.Start.Before(anotherEvent.Start) && e.End.After(anotherEvent.Start) {
		return true
	}

	if e.Start.Before(anotherEvent.End) && e.End.After(anotherEvent.End) {
		return true
	}

	return false
}

func (e Event) IsAValidEvent() bool {
	return e.Start.Before(e.End)
}
