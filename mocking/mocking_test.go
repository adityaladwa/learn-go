package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	CountDown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("Got '%s', Want '%s'", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("Not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
	}
}
