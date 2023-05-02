package main

import (
	"testing"
	"time"
)

func TestClock_GetTime(t *testing.T) {
	clock := Clock{}
	got := clock.GetTime()
	want := time.Now().Hour()
	if want != got {
		t.Errorf("Expected clock.GetTime() to be %v, got %v", want, got)
	}
}
