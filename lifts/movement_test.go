package lifts

import "testing"

func TestMovementString(t *testing.T) {
	cases := []struct {
		movement Movement
		s        string
	}{
		{Start, "Start"},
		{Stop, "Stop"},
	}

	for _, c := range cases {
		if c.movement.String() != c.s {
			t.Errorf("Expected %s as output to String, but got %s", c.s, c.movement.String())
		}
	}
}
