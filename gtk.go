// This file contains all the goey related code, including the gtk stuffs.
// It should be considered as part of ui, but was broken off to keep
// the file short to support readability.

package main

import (
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gtk"
	"golang.org/x/image/bmp"
	"image"
	"image/draw"
	"log"
)

var (
	mainWindow *goey.Window
)

func MakeBMP(digit *Time) error {
	var dg [3]*digit
	dg[0] = *digit.Hour
	dg[1] = *digit.Minute
	dg[2] = *digit.Second

	// need digit  to implement image.Image to encode to BMP.
	encode := bmp.Encode()

}

func MainWindow()
