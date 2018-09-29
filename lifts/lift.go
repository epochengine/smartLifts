package lifts

import (
	"sort"
	"time"
)

// Lift represents a lift and exposes functions to control it.
type Lift interface {
	AddDestination(floor int)
	Floor() int
	Destinations() []int
	ReportOn(chan int)
	Direction() Direction
}

// lift represents the internal state of a Lift.
type lift struct {
	floor        int
	speed        time.Duration
	destinations sort.IntSlice
	ch           chan int
}

// NewLift creates a Lift starting on the given floor.
func NewLift(startFloor int, speed time.Duration) Lift {
	return &lift{floor: startFloor, speed: speed}
}

// AddDestination adds the given destination to a lift.
func (l *lift) AddDestination(destination int) {
	travel := l.Direction() == Still
	insert := sort.SearchInts(l.destinations, destination)
	if insert == len(l.destinations) {
		l.destinations = append(l.destinations, destination)
	} else if l.destinations[insert] != destination {
		temp := make([]int, len(l.destinations)+1)
		copy(temp, l.destinations[:insert])
		temp[insert] = destination
		copy(temp[insert+1:], l.destinations[insert:])
		l.destinations = temp
	}

	if travel {
		go l.travel()
	}
}

func (l *lift) travel() {
	for len(l.destinations) != 0 {
		var diff int
		nextStop := l.destinations[0]
		if nextStop > l.floor {
			diff = 1
		} else {
			diff = -1
		}

		time.Sleep(l.speed)
		l.floor = l.floor + diff
		if l.floor == nextStop {
			l.destinations = l.destinations[1:]
		}
		if l.ch != nil {
			l.ch <- l.floor
		}
	}
}

// Floor returns the current floor of this lift.
func (l lift) Floor() int {
	return l.floor
}

// Destination returns the current destination list of this lift.
func (l lift) Destinations() []int {
	return l.destinations
}

// ReportOn instructs this lift to report its movements on the given channel.
func (l *lift) ReportOn(ch chan int) {
	l.ch = ch
}

// Direction returns the current movement direction of a lift.
func (l lift) Direction() Direction {
	switch {
	case len(l.destinations) == 0:
		return Still
	case l.destinations[0] < l.floor:
		return Down
	case l.destinations[0] > l.floor:
		return Up
	default:
		return Still
	}
}
