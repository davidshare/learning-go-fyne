package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	canvasApp := app.New()

	w := canvasApp.NewWindow("Icon: tutorial")
	w.Resize(fyne.NewSize(400, 400))
	iconX := widget.NewIcon(theme.CancelIcon())
	w.SetContent(iconX)
	w.ShowAndRun()
}
