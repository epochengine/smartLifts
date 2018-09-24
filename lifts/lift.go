package lifts

import (
	"math"
	"time"
)

// Lift represents a lift and exposes functions to control it.
type Lift interface {
	GoToFloor(floor int)
	Floor() int
	Destination() int
	ReportOn(chan int)
}

// lift respsents the internal state of a Lift.
type lift struct {
	floor       int
	speed       time.Duration
	destination int
	ch          chan int
}

// NewLift creates a Lift starting on the given floor.
func NewLift(startFloor int, speed time.Duration) Lift {
	return &lift{floor: startFloor, speed: speed}
}

// GoToFloor sends this lift to the given floor.
func (l *lift) GoToFloor(destination int) {
	l.destination = destination
	floors := int(math.Abs(float64(destination - l.floor)))
	time.Sleep(time.Duration(floors) * l.speed)
	l.floor = destination
	if l.ch != nil {
		l.ch <- destination
	}
}

// Floor returns the current floor of this lift.
func (l lift) Floor() int {
	return l.floor
}

// Destination returns the current destination of this lift.
func (l lift) Destination() int {
	return l.destination
}

// ReportOn instructs this lift to report its movements on the given channel.
func (l *lift) ReportOn(ch chan int) {
	l.ch = ch
}
