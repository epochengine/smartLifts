package lifts

import "testing"

func TestCreateLift(t *testing.T) {
	cases := []int{
		0,
		1,
	}

	for _, c := range cases {
		lift := NewLift(c)
		floor := lift.Floor()
		if floor != c {
			t.Errorf("Created lift on floor %d, got floor %d", c, floor)
		}
	}
}

func TestGoToFloor(t *testing.T) {
	lift := NewLift(0)
	lift.GoToFloor(1)
	if lift.Floor() != 1 {
		t.Errorf("Sent lift to floor %d, was actually on %d", 1, lift.Floor())
	}
}

func TestLiftReportsFloor(t *testing.T) {
	ch := make(chan (int))
	lift := NewLift(0)
	lift.ReportOn(ch)
	go lift.GoToFloor(1)
	report := <-ch
	if report != 1 {
		t.Errorf("Expected lift to report at floor %d, got %d", 1, report)
	}
}
