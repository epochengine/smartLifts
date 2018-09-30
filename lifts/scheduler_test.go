package lifts

import (
	"testing"
	"time"
)

func TestCallLift(t *testing.T) {
	cases := []int{
		1,
		2,
	}

	for _, c := range cases {
		liftScheduler := NewLiftScheduler()
		lift := NewLift(0, 0)
		ch := make(chan int)
		lift.ReportOn(ch)
		liftScheduler.RegisterLift(lift)
		lift, err := liftScheduler.CallLift(c)
		for i := 1; i <= c; i++ {
			<-ch
		}

		if err != nil {
			t.Errorf("Called lift to floor %d when one was available, but got error: %s", c, err)
			return
		}
		if lift.Floor() != c {
			t.Errorf("Called lift to floor %d, instead was at %d", c, lift.Floor())
		}
	}
}

func TestCallLiftSchedulesNearLift(t *testing.T) {
	liftScheduler := NewLiftScheduler()
	movingLift := NewLift(0, 1*time.Second)
	movingLift.AddDestination(10)
	farLift := NewLift(4, 1*time.Second)
	nearLift := NewLift(0, 1*time.Second)
	liftScheduler.RegisterLift(movingLift)
	liftScheduler.RegisterLift(farLift)
	liftScheduler.RegisterLift(nearLift)
	lift, err := liftScheduler.CallLift(1)

	if err != nil {
		t.Errorf("Expected the nearest lift to be scheduled, but instead got an error: %s", err)
		return
	}
	if lift != nearLift {
		t.Errorf("Expected the nearest lift to be scheduled, but instead got %v", lift)
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
