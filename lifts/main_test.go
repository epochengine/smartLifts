package lifts

import "testing"

func TestCallLift(t *testing.T) {
	cases := []int{
		1,
	}

	for _, c := range cases {
		lift := CallLift()
		if lift != c {
			t.Errorf("Called lift: want %d, got %d", c, lift)
		}
	}
}
