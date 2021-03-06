package lifts

// Direction represents a direction of travel (including none).
type Direction int

// A Direction can be Up, Down or Still.
const (
	Up Direction = iota
	Down
	Still
)

func (d Direction) String() string {
	switch d {
	case Up:
		return "Up"
	case Down:
		return "Down"
	default:
		return "Still"
	}
}
