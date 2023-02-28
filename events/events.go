package events

import (
	"fmt"
	"strings"
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

func BuildEventsFromLines(lines []string) (events []Event, err error) {
	const rfc2822 = "Mon Jan 02 15:04:05 -0700 2006"

	for _, line := range lines {
		dates := strings.Split(line, "\t")
		if len(dates) != 2 {
			return nil, fmt.Errorf("line must have 2 dates and has %d: %s", len(dates), line)
		}

		start, err := time.Parse(rfc2822, dates[0])
		if err != nil {
			return nil, err
		}
		end, err := time.Parse(rfc2822, dates[1])
		if err != nil {
			return nil, err
		}

		event := Event{
			Start: start,
			End:   end,
		}
		if !event.IsAValidEvent() {
			return nil, fmt.Errorf("given timelapse is not valid: %v", line)
		}

		events = append(events, event)
	}

	return events, nil
}
