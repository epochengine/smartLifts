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

func TestNewLiftScheduler(t *testing.T) {
	liftScheduler := NewLiftScheduler()
	lifts := liftScheduler.Lifts
	if len(lifts) != 4 {
		t.Errorf("Instantiated using NewLiftScheduler but length of lifts was %d, expected %d", len(lifts), 4)
	}
}
