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
	ReportFloorsOn(chan int)
	ReportMovementOn(chan Movement)
	Direction() Direction
}

// lift represents the internal state of a Lift.
type lift struct {
	floor        int
	speed        time.Duration
	destinations sort.IntSlice
	floorsCh     chan int
	movementCh   chan Movement
	travelling   bool
}

// NewLift creates a Lift starting on the given floor.
func NewLift(startFloor int, speed time.Duration) Lift {
	return &lift{floor: startFloor, speed: speed}
}

// AddDestination adds the given destination to a lift.
// Duplicate destinations will be silently discarded as the lift is already
// going there. If the given destination is the current floor, it will be
// ignored.
func (l *lift) AddDestination(destination int) {
	if destination == l.floor {
		return
	}

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

	if !l.travelling {
		l.travelling = true
		go l.travel()
	}
}

func (l *lift) travel() {
	if l.movementCh != nil {
		l.movementCh <- Start
	}

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
		if l.floorsCh != nil {
			l.floorsCh <- l.floor
		}
	}

	if l.movementCh != nil {
		l.movementCh <- Stop
	}
}

// Floor returns the current floor of a lift.
func (l lift) Floor() int {
	return l.floor
}

// Destination returns the current destination list of a lift.
func (l lift) Destinations() []int {
	return l.destinations
}

// ReportOn instructs a lift to report its movements on the given channel.
func (l *lift) ReportFloorsOn(ch chan int) {
	l.floorsCh = ch
}

// ReportMovementOn instructs a lift to report its start/stops on the
// given channel.
func (l *lift) ReportMovementOn(ch chan Movement) {
	l.movementCh = ch
}

// Direction returns the current movement direction of a lift.
func (l lift) Direction() Direction {
	switch {
	case len(l.destinations) == 0 || l.destinations[0] == l.floor:
		return Still
	case l.destinations[0] < l.floor:
		return Down
	default:
		return Up
	}
}
