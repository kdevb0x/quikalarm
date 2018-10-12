package main // import "github.com/kidoda/quikalarm"

import (
	"errors"
	"fmt"
	_ "image"
	co "image/color"
	"io"
	"log"
	"os"
	_ "strings"

	svg "github.com/ajstarks/svgo"
	cname "golang.org/x/image/colornames"
	// "github.com/mattn/gogtk/gtk"
	/* "github.com/fyne-io/fyne"
		"github.com/fyne-io/fyne/widget"
	    	"github.com/fyne-io/fyne/desktop"

	"github.com/fyne-io/fyne"
	*/)

var colormap = map[string]co.RGBA{
	"red":    cname.Red,
	"blue":   cname.Blue,
	"green":  cname.Green,
	"purple": cname.Violet,
	"white":  cname.White,
	"orange": cname.Orangered,
}

var COLOR string

type Digiter interface {
	SetNumber(n uint)
}

type Digit struct {
	Number   uint
	Out      io.Writer
	UpdateChan chan struct{}
	KillChan chan struct{}
	SVG      svg.SVG
	Color    co.RGBA
}

// UI is a Looper for displaying the Digital numbers.
type UI struct {
	T Time
}

// NewDigit creates a new Digit, default output is os.Stdout
// set Digit number with SetNumber method, set color with SetColor.
func NewDigit() *Digit {
	uchan := make(chan struct{}, 1)
	kchan := make(chan struct{}, 1)
	dig := &Digit{
		Out:      os.Stdout,
		UpdateChan uchan,
		KillChan: kchan,
	}
	return dig

}

// SetColor is used to set the color of a Digit in the Digit interface.
func SetColor(digi Digiter, color ...string) error {
	switch i := len(color); {
	case i > 1:
		err := errors.New("too many colors provided, please provide no more than one")
		return err

	case i == 1:
		c := color[0]
		col, ok := colormap[c]
		if !ok {
			err := errors.New("requested color not supported, --help for list")
			return err
		}
		n, ok := digi.(*Digit)
		n.Color = col
		digi.UpdateUI()

	}
	return nil
}

func (d *Digit) SetNumber(n uint) {
	d.Number = n

	sv := svg.New(os.Stdout)
	sv.Start(320, 600)
	var r, g, b, a = d.Color.RGBA()
	d.SVG.RGBA(int(r), int(g), int(b), float64(a))

	switch n {
	case 0:
		sv.Path("M22 19.5L42 0L278 0L298 19.5L278 39L42 39Z")
		sv.Path("M22 579.5L42 560L278 560L298 579.5L278 599L42 599Z")
		sv.Path("M 19.5 22L0 42L0 277.5L19.5 297.5L39 277.5L39 42Z")
		sv.Path("M299.5 22L280 42L280 277.5L299.5 297.5L319 277.5L319 42Z")
		sv.Path("M 19.5 302L0 322L0 557.5L19.5 577.5L39 557.5L39 322Z")
		sv.Path("M299.5 302L280 322L280 557.5L299.5 577.5L319 557.5L319 322Z")
		sv.End()
		d.SVG = *sv
		return

	}
}

// UpdateUI flushes any changes (such as changing the color) to the Output (io.Writer)
func (u *UI) Update() error {
	_, err := fmt.Println(u.T.Hour.SVG, u.T.Minute.SVG, u.T.Second.SVG)
	if err != nil {
		log.Printf("Error printing Digits to output: %s", err)
		return err
	}
	return nil
}

func (u *UI) Loop() {

}

type Application interface {
	MainWindow() *ui.Window
}


// func CreateMainWindow(app Application, title string) (*fyne.Window, error) {
// }

/*
print SVG "<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"320\" height=\"600\">\n";
     print SVG "<path fill=\"$colour\" d=\"M22 19.5L42 0L278 0L298 19.5L278 39L42 39Z\" />\n"			if grep { $_ == $Digit } (0,2,3,5,6,7,8,9);
     print SVG "<path fill=\"$colour\" d=\"M22 299.5L42 280L278 280L298 299.5L278 319L42 319Z\" />\n"		if grep { $_ == $Digit } (2,3,4,5,6,8,9);
     print SVG "<path fill=\"$colour\" d=\"M22 579.5L42 560L278 560L298 579.5L278 599L42 599Z\" />\n"		if grep { $_ == $Digit } (0,2,3,5,6,8,9);
     print SVG "<path fill=\"$colour\" d=\"M 19.5 22L0 42L0 277.5L19.5 297.5L39 277.5L39 42Z\" />\n"		if grep { $_ == $Digit } (0,4,5,6,7,8,9);
     print SVG "<path fill=\"$colour\" d=\"M299.5 22L280 42L280 277.5L299.5 297.5L319 277.5L319 42Z\" />\n"	if grep { $_ == $Digit } (0,1,2,3,4,7,8,9);
     print SVG "<path fill=\"$colour\" d=\"M 19.5 302L0 322L0 557.5L19.5 577.5L39 557.5L39 322Z\" />\n"		if grep { $_ == $Digit } (0,2,6,8);
     print SVG "<path fill=\"$colour\" d=\"M299.5 302L280 322L280 557.5L299.5 577.5L319 557.5L319 322Z\" />\n"	if grep { $_ == $Digit } (0,1,3,4,5,6,7,8,9);

*/
