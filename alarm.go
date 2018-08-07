package main

import (
	_ "fmt"
	_ "io/ioutil"
	"time"

	_ "github.com/faiface/beep"
	_ "github.com/faiface/beep/speaker"
	_ "github.com/faiface/beep/wav"
	_ "github.com/go-bindata/go-bindata"
)

type clock struct {
	CurrentTime time.Time
	WakeTime    time.Time
}

func NewClock(ct time.Time) clock {
	c := &clock{}
	c.CurrentTime = ct
	return *c
}

func (c *clock) SetWakeTime(t time.Time, byDur ...time.Duration) error {
	if len(byDur) > 0 {
		c.WakeTime = c.CurrentTime + byDur[0]
	}
}
