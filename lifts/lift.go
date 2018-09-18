package lifts

// Lift represents a lift and its state.
type Lift struct {
	Floor int
}

// GoToFloor will send this lift to the given floor.
func (l *Lift) GoToFloor(floor int) {
	l.Floor = floor
}
