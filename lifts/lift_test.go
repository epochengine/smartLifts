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
		lift.AddDestination(c.destination)
		destinations := lift.Destinations()
		if len(destinations) != 1 && destinations[0] != c.destination {
			t.Errorf("Sent lift to floor %d, but its destination was actually %d", c.destination, destinations[0])
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
		go lift.AddDestination(c)

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
		go lift.AddDestination(c.floors)
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

func TestAddDestination(t *testing.T) {
	lift := NewLift(0, 1*time.Second)
	lift.AddDestination(1)
	lift.AddDestination(4)
	lift.AddDestination(2)
	expected := []int{1, 2, 4}
	destinations := lift.Destinations()

	if len(expected) != len(destinations) {
		t.Errorf("lift.Destinations() should be length %d, but was instead %d", len(expected), len(destinations))
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != destinations[i] {
			t.Errorf("Expected destination at index %d is %d but got %d", i, expected[i], destinations[i])
		}
	}
}

func TestDirection(t *testing.T) {
	cases := []struct {
		start       int
		destination int
		direction   Direction
	}{
		{0, 2, Up},
		{2, 0, Down},
		{0, 0, Still},
	}

	for _, c := range cases {
		lift := NewLift(c.start, 10*time.Millisecond)
		lift.AddDestination(c.destination)
		if lift.Direction() != c.direction {
			t.Errorf("Sent lift from %d to %d and expected direction %s, but got %s", c.start, c.destination, c.direction.String(), lift.Direction())
		}
	}
}

func TestDirectionAfterMovement(t *testing.T) {
	lift := NewLift(0, 10*time.Millisecond)
	ch := make(chan int)
	lift.ReportOn(ch)
	lift.AddDestination(1)
	<-ch
	if lift.Direction() != Still {
		t.Errorf("Expected lift to be Still after moving floors, but instead its direction was %s", lift.Direction())
	}
}
