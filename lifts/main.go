package lifts

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
// It returns the lift number that has been assigned.
func CallLift() int {
	return 1
}

// RegisterLift adds a Lift to the system, available for scehduling.
func (ls liftScheduler) RegisterLift(lift Lift) {
	ls.Lifts[lift] = struct{}{}
}
