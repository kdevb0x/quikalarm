package main // import "github.com/kidoda/quikalarm"

import (
	"errors"
	"fmt"
	_ "image"
	co "image/color"
	"io"
	_ "log"
	"os"
	_ "strings"

	svg "github.com/ajstarks/svgo"
	// "github.com/gotk3/gotk3/gtk"
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

type Digit interface {
	SetNumber(n uint)
	UpdateUI()
}

type digit struct {
	Number     uint
	Out        io.Writer
	UpdateChan chan struct{}
	KillChan   chan struct{}
	SVG        svg.SVG
	Color      co.RGBA
}

// NewDigit creates a new digit, default output is os.Stdout
// set digit number with SetNumber method, set color with SetColor.
func NewDigit() *digit {
	uchan := make(chan struct{}, 1)
	kchan := make(chan struct{}, 1)
	dig := &digit{
		Out:        os.Stdout,
		UpdateChan: uchan,
		KillChan:   kchan,
	}
	return dig

}

// SetColor is used to set the color of a digit in the Digit interface.
func SetColor(digi Digit, color ...string) error {
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
		n, ok := digi.(*digit)
		n.Color = col
		digi.UpdateUI()

	}
	return nil
}

func (d *digit) SetNumber(n uint) {
	d.Number = n

	// number := digit{Number: 1, Out: os.Stdout, UpdateChan: Uchan, KillChan: Kchan, SVG: nil,}
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
func (d *digit) UpdateUI() {
	fmt.Println(d.SVG)
}

type Application interface {
	// fyne.App
}

// func CreateMainWindow(app Application, title string) (*fyne.Window, error) {
// }

/*
print SVG "<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"320\" height=\"600\">\n";
     print SVG "<path fill=\"$colour\" d=\"M22 19.5L42 0L278 0L298 19.5L278 39L42 39Z\" />\n"			if grep { $_ == $digit } (0,2,3,5,6,7,8,9);
     print SVG "<path fill=\"$colour\" d=\"M22 299.5L42 280L278 280L298 299.5L278 319L42 319Z\" />\n"		if grep { $_ == $digit } (2,3,4,5,6,8,9);
     print SVG "<path fill=\"$colour\" d=\"M22 579.5L42 560L278 560L298 579.5L278 599L42 599Z\" />\n"		if grep { $_ == $digit } (0,2,3,5,6,8,9);
     print SVG "<path fill=\"$colour\" d=\"M 19.5 22L0 42L0 277.5L19.5 297.5L39 277.5L39 42Z\" />\n"		if grep { $_ == $digit } (0,4,5,6,7,8,9);
     print SVG "<path fill=\"$colour\" d=\"M299.5 22L280 42L280 277.5L299.5 297.5L319 277.5L319 42Z\" />\n"	if grep { $_ == $digit } (0,1,2,3,4,7,8,9);
     print SVG "<path fill=\"$colour\" d=\"M 19.5 302L0 322L0 557.5L19.5 577.5L39 557.5L39 322Z\" />\n"		if grep { $_ == $digit } (0,2,6,8);
     print SVG "<path fill=\"$colour\" d=\"M299.5 302L280 322L280 557.5L299.5 577.5L319 557.5L319 322Z\" />\n"	if grep { $_ == $digit } (0,1,3,4,5,6,7,8,9);

*/
