package main // import "github.com/kidoda/quikalarm"

import (
	"io"
	"time"
)

type Time struct {
	TimeChan []time.Ticker
	Hour     *Digit
	Minute   *Digit
	Second   *Digit
}

// Looper monitors an UpdateChan for members that have signaled a change.
// Once all of the adjustable threshold parameters have been met, it batches the changes
// and then triggers the Updater, which handles the execution.
type Looper interface {
	Loop()
}

// Updater handles the execution of the batched gui updates handed off to it from a Looper.
type Updater interface {
	Update()
}

// *sidenote: This mechanism was chosen over calllbacks because of the frequency
// of the updates, it didnt make sense to me to use callbacks for something that
// is guarenteed to run at least 60 times per minute constantly.
// ~(kdev) TODO: Needs battle testing. If this 2 part idiom doesn't facilitate a fluid
// and enjoyable development experience, scrap it; possibly for the common single
// update loop idiom.

// Loop method of Time tracks time.Ticker and readies the correct Digits for ui.
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
