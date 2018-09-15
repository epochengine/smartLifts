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
