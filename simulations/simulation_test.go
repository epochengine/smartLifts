package simulations

import (
	"testing"

	"epochengine/liftscheduler/lifts"
)

func TestSimulateProvidesEventList(t *testing.T) {
	simulation := Simulation{}

	var events [3]Event
	events[0] = Event{1, lifts.Down}
	events[1] = Event{2, lifts.Up}
	events[2] = Event{0, lifts.Down}

	for i := 0; i < len(events); i++ {
		simulation.AddEvent(events[i])
	}

	count := 0
	receiver := func(e Event) {
		if e != events[count] {
			t.Errorf("Got incorrect event. Expected %v but got %v", events[count], e)
		}
		count++
	}

	simulation.Simulate(receiver)
}
