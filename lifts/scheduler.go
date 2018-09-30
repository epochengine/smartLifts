package lifts

import (
	"errors"
)

// LiftScheduler schedules a set of Lifts.
type liftScheduler struct {
	Lifts map[Lift]struct{}
}

// NewLiftScheduler creates a liftScheduler.
// It will have an empty set of Lifts.
func NewLiftScheduler() liftScheduler {
	lifts := make(map[Lift]struct{})
	return liftScheduler{lifts}
}

// CallLift requests a lift.
// It returns the lift that has been assigned.
func (ls liftScheduler) CallLift(floor int, direction Direction) (l Lift, err error) {
	var bestLift Lift
	for lift := range ls.Lifts {
		if lift.Direction() == Still && liftIsCloser(bestLift, lift, floor) {
			bestLift = lift
		} else if lift.Direction() == direction && liftIsCloser(bestLift, lift, floor) {
			bestLift = lift
		}
	}

	if bestLift != nil {
		bestLift.AddDestination(floor)
		return bestLift, nil
	}

	return nil, errors.New("no lift available to call")
}

func liftIsCloser(bestLift Lift, candidateLift Lift, targetFloor int) bool {
	if bestLift == nil {
		return true
	}

	bestLiftDiff := bestLift.Floor() - targetFloor
	if bestLiftDiff < 0 {
		bestLiftDiff = -bestLiftDiff
	}
	candidateLiftDiff := candidateLift.Floor() - targetFloor
	if candidateLiftDiff < 0 {
		candidateLiftDiff = -candidateLiftDiff
	}

	return candidateLiftDiff < bestLiftDiff
}

// RegisterLift adds a Lift to the system, available for scheduling.
func (ls liftScheduler) RegisterLift(lift Lift) {
	ls.Lifts[lift] = struct{}{}
}
