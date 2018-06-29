package main

import (
	"fmt"
	"time"
	""
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)
// Clock ticks by the second.
type clock struct {
	startTime time.Time
	ticker time.Ticker
}

type alarm struct {
	clock clock
	wakeTime time.Time
}

type Alarm interface {
	Set(time
}

func newClock() clock {
	clock := clock{
		startTime: time.Now()
		ticker: time.NewTicker(time.Second)
	}
	return clock

}

var
