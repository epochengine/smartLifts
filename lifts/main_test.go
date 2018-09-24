package lifts

import "testing"

func TestCallLift(t *testing.T) {
	cases := []int{
		1,
		2,
	}

	for _, c := range cases {
		liftScheduler := NewLiftScheduler()
		liftScheduler.RegisterLift(NewLift(0, 0))
		lift, err := liftScheduler.CallLift(c)
		if lift.Floor() != c {
			t.Errorf("Called lift to floor %d, instead was at %d", c, lift.Floor())
		}
		if err != nil {
			t.Errorf("Called lift to floor %d when one was available, but got error: %s", c, err)
		}
	}
}

func TestCallLiftError(t *testing.T) {
	liftScheduler := NewLiftScheduler()
	_, err := liftScheduler.CallLift(1)
	if err == nil {
		t.Error("Called lift when none was available. Expected error but got none.")
	}
}

func TestRegisterLift(t *testing.T) {
	liftScheduler := NewLiftScheduler()
	lift := NewLift(0, 0)
	liftScheduler.RegisterLift(lift)
	lifts := liftScheduler.Lifts
	_, ok := lifts[lift]
	if !ok {
		t.Error("Called RegisterLift but lift was not found in Lifts")
	}
}
