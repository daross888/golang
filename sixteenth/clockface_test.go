package sixteenth_test

import (
	"learning/sixteenth"
	"testing"
	"time"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := sixteenth.Point{X: 150, Y: 150 - 90}
	got := sixteenth.SecondHand(tm)

	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestSecondHandAt30Seconds(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

	want := sixteenth.Point{X: 150, Y: 150 + 90}
	got := sixteenth.SecondHand(tm)

	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}
