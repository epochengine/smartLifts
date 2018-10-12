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
		lift.ReportFloorsOn(ch)
		liftScheduler.RegisterLift(lift)
		lift, err := liftScheduler.CallLift(c, Down)
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
	lift, err := liftScheduler.CallLift(1, Down)

	if err != nil {
		t.Errorf("Expected the nearest lift to be scheduled, but instead got an error: %s", err)
		return
	}
	if lift != nearLift {
		t.Errorf("Expected the nearest lift to be scheduled, but instead got %v", lift)
	}
}

func TestCallLiftSchedulesLiftHeadingTowardsCall(t *testing.T) {
	cases := []struct {
		towardsLiftStart int
		towardsLiftDest  int
		awayLiftStart    int
		awayLiftDest     int
		callFloor        int
		direction        Direction
	}{
		{5, 0, 2, 10, 1, Down},
		{2, 10, 5, 0, 4, Up},
	}

	for _, c := range cases {
		liftScheduler := NewLiftScheduler()
		awayLift := NewLift(c.awayLiftStart, 1*time.Second)
		towardsLift := NewLift(c.towardsLiftStart, 1*time.Second)
		liftScheduler.RegisterLift(awayLift)
		liftScheduler.RegisterLift(towardsLift)
		awayLift.AddDestination(c.awayLiftDest)
		towardsLift.AddDestination(c.towardsLiftDest)
		lift, err := liftScheduler.CallLift(1, c.direction)

		if err != nil {
			t.Errorf("Expected the lift heading towards call to be scheduled, but instead got an error: %s", err)
			return
		}
		if lift != towardsLift {
			t.Errorf("Expected the lift heading towards call to be scheduled, but instead got %v", lift)
		}
	}
}

func TestCallLiftError(t *testing.T) {
	liftScheduler := NewLiftScheduler()
	_, err := liftScheduler.CallLift(1, Down)
	if err == nil {
		t.Error("Called lift when none was available. Expected error but got none.")
	}
}

func TestLiftIsCloser(t *testing.T) {
	cases := []struct {
		bestLiftFloor      int
		candidateLiftFloor int
		targetFloor        int
		expected           bool
	}{
		{5, 1, 2, true},
		{5, 4, 2, true},
		{5, 1, 3, false},
		{1, 5, 4, true},
		{-1, 0, -2, false},
		{0, -1, -2, true},
	}

	for _, c := range cases {
		bestLift := NewLift(c.bestLiftFloor, 0)
		candidateLift := NewLift(c.candidateLiftFloor, 0)
		result := liftIsCloser(bestLift, candidateLift, c.targetFloor)
		if result != c.expected {
			t.Errorf("Expected liftIsCloser to return %v but got %v. Best lift floor was %d, candidate lift floor was %d, and target floor was %d",
				c.expected, result, c.bestLiftFloor, c.candidateLiftFloor, c.targetFloor)
		}
	}
}

func TestLiftIsCloserNilBest(t *testing.T) {
	candidateLift := NewLift(100, 0)
	result := liftIsCloser(nil, candidateLift, -100)
	if !result {
		t.Errorf("Expected liftIsCloser to always return true if no best lift is supplied, but got %v", result)
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
