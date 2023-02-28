package eventsloader

import (
	"fmt"
	"strings"
	"time"

	"github.com/EdigiraldoBF/double-booked/events"
	"github.com/EdigiraldoBF/double-booked/utils/file"
)

type EventsFromTxtLoader struct {
	filepath string
}

func NewEventsFromTxtLoader(filepath string) EventsFromTxtLoader {
	return EventsFromTxtLoader{
		filepath: filepath,
	}
}

func (l *EventsFromTxtLoader) Load() (events.Events, error) {
	lines, err := file.GetLines(l.filepath)
	if err != nil {
		return events.Events{}, fmt.Errorf("failed loading file: %v", err)
	}

	evs, err := BuildEventsFromLines(lines)
	if err != nil {
		return events.Events{}, fmt.Errorf("failed building events: %v", err)
	}

	return evs, nil
}

func BuildEventsFromLines(lines []string) (evs events.Events, err error) {
	const rfc2822 = "Mon Jan 02 15:04:05 -0700 2006"

	for _, line := range lines {
		dates := strings.Split(line, "\t")
		if len(dates) != 2 {
			return events.Events{}, fmt.Errorf("line must have 2 dates and has %d: %s", len(dates), line)
		}

		start, err := time.Parse(rfc2822, dates[0])
		if err != nil {
			return events.Events{}, err
		}
		end, err := time.Parse(rfc2822, dates[1])
		if err != nil {
			return events.Events{}, err
		}

		event := events.Event{
			Start: start,
			End:   end,
		}
		if !event.IsAValidEvent() {
			return events.Events{}, fmt.Errorf("given timelapse is not valid: %v", line)
		}

		evs.Events = append(evs.Events, event)
	}

	return evs, nil
}
