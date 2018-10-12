package lifts

// Movement exposes when a lift starts or stops moving.
type Movement int

// Movement is expressed only as Start or Stop.
const (
	Start Movement = iota
	Stop
)

func (m Movement) String() string {
	switch m {
	case Start:
		return "Start"
	default:
		return "Stop"
	}
}
