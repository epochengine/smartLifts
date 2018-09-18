package lifts

import "testing"

func TestCreateLift(t *testing.T) {
	cases := []int{
		0,
		1,
	}

	for _, c := range cases {
		lift := Lift{c}
		floor := lift.Floor
		if floor != c {
			t.Errorf("Created lift on floor %d, got floor %d", c, floor)
		}
	}
}

func TestGoToFloor(t *testing.T) {
	lift := Lift{0}
	lift.GoToFloor(1)
	if lift.Floor != 1 {
		t.Errorf("Sent lift to floor %d, was actually on %d", 1, lift.Floor)
	}
}
