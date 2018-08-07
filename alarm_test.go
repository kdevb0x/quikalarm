package main

import (
	_ "io/ioutil"
	"testing"
	"time"

	_ "github.com/faiface/beep"
	_ "github.com/faiface/beep/speaker"
	_ "github.com/faiface/beep/wav"
	_ "github.com/go-bindata/go-bindata"
)

func TestNewClock(t *testing.T) {
	clock := NewClock(time.Now())
	time.Sleep(3 * time.Second)
	if n := time.Now(); n != clock.CurrentTime {
		t.Fail("clock's current time incorrect; != time.Now()")
	}
}

func TestSetWakeTime(t *testing.T) {
	clock := NewClock(time.Now())
	clock.SetWakeTime(5 * time.Hour)
	if clock.WakeTime != time.Now()+(5*time.Hour) {
		t.Fail("failed to set time")
	}
}
