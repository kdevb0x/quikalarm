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
	cl := NewClock(time.Now())

	var lapse time.Duration = 3 * time.Second
	time.Sleep(lapse)

	if n := cl.CreationTime.Sub(time.Now()); n != lapse {
		t.Fail()
	}
}

func TestSetWakeTime(t *testing.T) {
	clock := NewClock(time.Now())
	clock.SetWakeTime(timeOrDuration(5*time.Hour), true)
	if clock.WakeTime != time.Now().Add(5*time.Hour) {
		t.Fail()
	}
}
