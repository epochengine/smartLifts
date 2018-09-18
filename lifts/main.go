package lifts

// LiftScheduler schedules a set of Lifts.
type liftScheduler struct {
	Lifts []Lift
}

// NewLiftScheduler creates a liftScheduler.
// It will have an empty but allocated slice of Lifts.
func NewLiftScheduler() liftScheduler {
	lifts := make([]Lift, 4)
	return liftScheduler{lifts}
}

// CallLift requests a lift.
// It returns the lift number that has been assigned.
func CallLift() int {
	return 1
}

// RegisterLift adds a Lift to the system, available for scehduling.
func (ls liftScheduler) RegisterLift(lift Lift) {
	ls.Lifts = append(ls.Lifts, lift)
}
