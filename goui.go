package main

import (
	"os"
	"runtime"
	"strconv"

	"github.com/gotk3/gotk3/gtk"

	"github.com/andlabs/ui"
)

type ApplicationUI struct {
	mainWindow *ui.Window
	buttons    []ui.Button
}

type GUI interface {
	Show()
	NewMainWindow() *gtk.
}

func MainWindow() *ui.Window {
	f := func() (*ui.Window, error) {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()

		window := ui.Main(func() {
			win := ui.NewWindow("QuikAlarm", 360, 240, true)
			setWakeTime := ui.NewButton("SET WAKE TIME")
			picker := ui.NewTimePicker()
			picker.SetTime(time.Now)

		})
	}
}
