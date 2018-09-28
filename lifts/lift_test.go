package lifts

import (
	"testing"
	"time"
)

func TestCreateLift(t *testing.T) {
	cases := []int{
		0,
		1,
	}

	for _, c := range cases {
		lift := NewLift(c, 0)
		floor := lift.Floor()
		if floor != c {
			t.Errorf("Created lift on floor %d, got floor %d", c, floor)
		}
	}
}

func TestGoToFloor(t *testing.T) {
	cases := []struct {
		start       int
		destination int
	}{
		{0, 2},
		{3, 1},
	}

	for _, c := range cases {
		lift := NewLift(c.start, 50*time.Millisecond)
		lift.GoToFloor(c.destination)
		if lift.Destination() != c.destination {
			t.Errorf("Sent lift to floor %d, but its destination was actually %d", c.destination, lift.Destination())
		}
	}
}

func TestLiftReportsFloor(t *testing.T) {
	cases := []int{
		1,
		2,
	}

	for _, c := range cases {
		ch := make(chan int)
		lift := NewLift(0, 20*time.Millisecond)
		lift.ReportOn(ch)
		go lift.GoToFloor(c)

		for i := 1; i <= c; i++ {
			report := <-ch
			if report != i {
				t.Errorf("Expected lift to report at floor %d, got %d", i, report)
			}
		}
	}
}

func TestLiftMovementSpeed(t *testing.T) {
	cases := []struct {
		floors int
		speed  time.Duration
	}{
		{1, 80 * time.Millisecond},
		{2, 50 * time.Millisecond},
	}

	for _, c := range cases {
		lift := NewLift(0, c.speed)
		ch := make(chan int)
		lift.ReportOn(ch)
		start := time.Now()
		go lift.GoToFloor(c.floors)
		for i := 1; i <= c.floors; i++ {
			<-ch
		}

		duration := time.Since(start)
		expected := c.speed * time.Duration(c.floors)
		if duration < expected {
			t.Errorf("Expected lift to take at least %s to travel %d floors at speed %s per floor, instead took %s", expected, c.floors, c.speed, time.Duration(duration))
		}
	}
}
