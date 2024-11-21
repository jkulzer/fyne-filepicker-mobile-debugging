package main

import (
	"fmt"
	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
)

func main() {
	a := app.NewWithID("1")
	w := a.NewWindow("testwindow")
	filePicker := dialog.NewFileOpen(func(f fyne.URIReadCloser, err error) {
		fmt.Println("file selected")
	}, w)
	filePicker.Show()

	w.ShowAndRun()
}
