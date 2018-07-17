package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

func TestLoadBuzzer(t *testing.T) {
	buzz := loadBuzzer()
	var streamLen = buzz.stream.Len()
	if streamLen !> 0 {
		t.Fail("Failed to loadBuzzer, streamer length is 0")
	}
}

func TestPlayBuzzer(t *testing.T) {
	buzz := loadBuzzer()
	go func() {
		sr := beep.SampleRate(44, 100)
		initPlaybackBuff := speaker.Init(sr, sr.N(time.Millisecond*50))
	}()
	loopStream := beep.Loop(-1, buzz.stream)
	speaker.Play(loopStream)
}
