package main // import "github.com/kidoda/quikalarm"

import (
	"io"
	"time"
)

type Time struct {
	TimeChan []time.Ticker
	Hour     *digit
	Minute   *digit
	Second   *digit
}

type Looper interface {
	Loop()
}

// Time loop tracks time.Ticker and readies the correct digits for ui.
func (t *Time) Loop() {
	for {
		select {
		case hr := <-t.TimeChan[0].C:
			t.Hour.SetNumber(uint(hr.Hour()))
		case mn := <-t.TimeChan[1].C:
			t.Minute.SetNumber(uint(mn.Minute()))
		case sc := <-t.TimeChan[2].C:
			t.Second.SetNumber(uint(sc.Second()))
		}
	}
}

type OutputWriter interface {
	io.Writer
}
