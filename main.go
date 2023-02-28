package main

import (
	"fmt"
	"log"

	"github.com/EdigiraldoBF/double-booked/events"
	"github.com/EdigiraldoBF/double-booked/eventsloader"
)

func main() {
	filepath := "./date-ranges.txt"

	TxtLoader := eventsloader.NewEventsFromTxtLoader(filepath)
	eventsloader.SetLoader(&TxtLoader)

	evs, err := eventsloader.Load()
	if err != nil {
		log.Fatal(err)
	}

	overlappingEvents := GetOverlappingEvents(evs.Events)

	if len(overlappingEvents) != 0 {
		fmt.Printf("The following are the overlapping events in the provided list:\n\n")
	} else {
		fmt.Printf("There are no overlapping events in the provided list\n")
	}
	for _, overlappedPair := range overlappingEvents {
		fmt.Printf("%+v\n", overlappedPair)
	}
}

func GetOverlappingEvents(evs []events.Event) (overlappingEvents [][]events.Event) {
	for i := range evs {
		eventA := evs[i]
		for j := i + 1; j < len(evs); j++ {
			eventB := evs[j]
			if eventA.OverlapsWith(eventB) {
				overlappedPair := []events.Event{eventA, eventB}
				overlappingEvents = append(overlappingEvents, overlappedPair)
			}
		}
	}

	return overlappingEvents
}
