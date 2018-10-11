package simulations

import "epochengine/liftscheduler/lifts"

// A Simulation represents a set of Events to be run against a system.
// Events are played back in the order they're added.
type Simulation struct {
	events []Event
}

// Event defines an individual event to be run in a simulation.
type Event struct {
	floor     int
	direction lifts.Direction
}

// AddEvent adds an event to a simulation's list.
func (s *Simulation) AddEvent(e Event) {
	s.events = append(s.events, e)
}

// Simulate plays back the events in a simulation's list.
func (s *Simulation) Simulate(receiver func(Event)) {
	for _, e := range s.events {
		receiver(e)
	}
}
