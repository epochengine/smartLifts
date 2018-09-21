package lifts

// Lift represents a lift and exposes functions to control it.
type Lift interface {
	GoToFloor(floor int)
	Floor() int
	ReportOn(chan int)
}

// lift respsents the internal state of a Lift.
type lift struct {
	floor int
	ch    chan int
}

// NewLift creates a Lift starting on the given floor.
func NewLift(startFloor int) Lift {
	return &lift{floor: startFloor}
}

// GoToFloor sends this lift to the given floor.
func (l *lift) GoToFloor(floor int) {
	l.floor = floor
	if l.ch != nil {
		l.ch <- floor
	}
}

// Floor returns the current floor of this lift.
func (l lift) Floor() int {
	return l.floor
}

// ReportOn instructs this lift to report its movements on the given channel.
func (l *lift) ReportOn(ch chan int) {
	l.ch = ch
}
