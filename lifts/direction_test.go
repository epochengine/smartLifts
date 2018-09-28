package lifts

import "testing"

func TestDirectionString(t *testing.T) {
	cases := []struct {
		direction Direction
		s         string
	}{
		{Up, "Up"},
		{Down, "Down"},
		{Still, "Still"},
	}

	for _, c := range cases {
		if c.direction.String() != c.s {
			t.Errorf("Expected %s as output to String, but got %s", c.s, c.direction.String())
		}
	}
}
