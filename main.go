package main // import "github.com/kidoda/quikalarm"

import (
	"log"
	"os"

	"github.com/fyne-io/fyne/desktop"
)

func main() {

	app := desktop.NewApp()
	mainWindow, err := CreateMainWindow(app, "QuikAlarm")
	if err != nil {
		log.Fatalf("Fatal error: %s", err)
	}
}
