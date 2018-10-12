package main // import "toy"

import (
	"fmt"
	"runtime"

	"github.com/andlabs/ui"
)

func main() {
	MainWindow()

}

func MainWindow() error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	err := ui.Main(func() {
		win := ui.NewWindow("test", 800, 600, true)
		button := ui.NewButton("CLICK ME!")
		blank := ui.NewProgressBar()
		win.SetChild(button)
		button.OnClicked(func(*ui.Button) {
			count := 0
			fmt.Printf("clicked")
			blank.SetValue(count)
			count += 1
		})
		win.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true

		})
		win.Show()
	})
	if err != nil {
		panic(err)
	}
	return nil
}
