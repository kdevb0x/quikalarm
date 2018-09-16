package main // import "github.com/kidoda/quikalarm"

import (
	"io"
	"log"
	"time"

	svg "github.com/ajstarks/svgo"
	"github.com/gotk3/gotk3/gtk"
	// "github.com/mattn/gogtk/gtk"
	/* "github.com/fyne-io/fyne"
	"github.com/fyne-io/fyne/widget"
    	"github.com/fyne-io/fyne/desktop"
	*/
)

type Digit interface {
	UpdateCanvas(canvas *svg.SVG)
	PrintSVG(out io.Writer)
}

type digit stuct {
	Number uint8
	Out io.Writer
	UpdateChan chan struct{}
	KillChan chan struct{}
	SVG svg.SVG
}

type Application interface {
	// fyne.App
}
func CreateMainWindow(app Application, title string) (*fyne.Window, error) {
}



func (d *digit) UpdateCanvas(canvas)
