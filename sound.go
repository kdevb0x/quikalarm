package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/faiface/beep"
	_ "github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

type buzzer struct {
	assetData io.Reader
	stream    beep.StreamSeekCloser
}

// LoadBuzzer initializes the sound from sound asset
func loadBuzzer() *buzzer {
	sound, err := alarmWav()
	var assetReader = bytes.NewReader(sound.bytes)
	if err != nil {
		fmt.Printf("Couldn't load alarmWavBytes for new buzzer. ERROR: %s \n", err)
		return nil
	}
	var assetReadCloser = ioutil.NopCloser(assetReader)
	stream, _, err := wav.Decode(assetReadCloser)
	if err != nil {
		errors.New("Unable to decode wav")
	}
	buz := &buzzer{
		assetData: assetReader,
		stream:    stream,
	}
	return buz
}

/*
// TODO: Possibly delete as this is the same as loadBuzzer
// loadSound reads in the static soundfile from binary and returns a pointer
func loadSound() (*Buzzer, error) {
	sound, err := alarmWavBytes()
	if err != nil {
		return nil, err
	}
	var bz buzzer
	bz.assetData = sound
	bytes.NewReader(sound)
	bz.Stream = wav.Decode

	var streamer beep.Streamer
	sndBuffer := beep.NewBuffer(beep.Format{408442, 1, 2})

}
*/
