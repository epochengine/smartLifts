package lifts

import "errors"

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
func (ls liftScheduler) CallLift(floor int) (l Lift, err error) {
	for lift := range ls.Lifts {
		lift.GoToFloor(floor)
		return lift, nil
	}

	return nil, errors.New("no lift available to call")
}

// RegisterLift adds a Lift to the system, available for scehduling.
func (ls liftScheduler) RegisterLift(lift Lift) {
	ls.Lifts[lift] = struct{}{}
}
