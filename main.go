package main

import (
	"fmt"
	"log"

	"github.com/EdigiraldoBF/double-booked/events"
	"github.com/EdigiraldoBF/double-booked/utils/file"
)

func main() {
	lines, err := file.GetLines("./date-ranges.txt")

	if err != nil {
		log.Fatalf("failed loading file: %v", err)
	}

	evs, err := events.BuildEventsFromLines(lines)
	if err != nil {
		log.Fatalf("failed building events: %v", err)
	}

	overlappingEvents := GetOverlappingEvents(evs)

	if len(overlappingEvents) != 0 {
		fmt.Println("The following are the overlapping events in the provided list:")
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
