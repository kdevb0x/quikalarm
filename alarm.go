package main

import (
	"fmt"
	_ "io/ioutil"
	"log"
	"time"

	_ "github.com/faiface/beep"
	_ "github.com/faiface/beep/speaker"
	_ "github.com/faiface/beep/wav"
	_ "github.com/go-bindata/go-bindata"
)

type clock struct {
	CreationTime time.Time
	CurrentTime  time.Time
	WakeTime     time.Time
}

// timeOrDuration makes it possible to set Waketime by time.Time
// or time.Duration depending on the argument type given when called.
type timeOrDuration interface {
	String() string
	Now() pkgTime
}

type pkgTime struct {
	time.Time
}

func (pt pkgTime) String() string {
	return fmt.Sprintf("%#v", pt)
}

// NewClock returns a new instance of clock.
func NewClock() Clock {
	now := pkgTime.Now()
	c := &clock{}
	c.CreationTime = now
	return c
}

// Clock is a wrapper for concrete type clock.
type Clock interface {
	SetWakeTime(timeOrDuration, bool) error
}

func (c *clock) SetWakeTime(tm timeOrDuration, isDuration bool) error {
	if isDuration == true {
		parsedDur, err := time.ParseDuration(tm.String())
		if err != nil {
			log.Println(err)
			return err
		}
		c.WakeTime = c.CurrentTime.Add(parsedDur)
	}
	return nil

}
