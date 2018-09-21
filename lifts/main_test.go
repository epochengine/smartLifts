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

func TestRegisterLift(t *testing.T) {
	liftScheduler := NewLiftScheduler()
	lift := NewLift(0)
	liftScheduler.RegisterLift(lift)
	lifts := liftScheduler.Lifts
	_, ok := lifts[lift]
	if !ok {
		t.Error("Called RegisterLift but lift was not found in Lifts")
	}
}
