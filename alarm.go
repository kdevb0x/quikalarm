package main // import "github.com/kidoda/quikalarm"

import (
	_ "errors"
	"time"

	_ "github.com/faiface/beep"
	_ "github.com/faiface/beep/speaker"
	_ "github.com/faiface/beep/wav"
)

var (
	now        = time.Now()   // A global now just for reference
	alarmSound = loadBuzzer() // TODO: Git rid of this global and call locally in an alarm field.
	mainClock  = NewClock()
)

// clock is the time keeping part of an alarm.
type clock struct {
	startTime  time.Time
	currentTimeChan func() <-chan string {
		ch := make(chan string, 1)
	}
	currentTime time.Time
}

func (a *alarm) SetWakeTime(newWakeTime time.Time) {
	if after := a.currentTime
}

func (a alarm) GetWakeTime() string {
	return a.wakeTime.String()
}
// alarm is the complete object with time, buzzer, and wake time.
type alarm struct {
	clock      *clock
	wakeTime   time.Time
	alarmSound buzzer
}
// TODO: Need to rethink this approach, combine creations?
func NewAlarm() *alarm {
	anAlarm := &alarm{
		clock: NewClock(),
		alarmSound: loadBuzzer(),
	}
	return anAlarm
}

// NewClock instantiates a new clock with current time values and returns a pointer
func NewClock() *clock {
	clock := &clock{
		startTime: time.Now(),
		ticker:    time.NewTicker(time.Second * 1),
	}
	go func() {
		for t := range clock.ticker.C {
			clock.currentTimeString = t.String()
		}
	}()

	return clock

}
