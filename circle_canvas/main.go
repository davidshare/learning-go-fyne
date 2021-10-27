package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	canvasApp := app.New()

	w := canvasApp.NewWindow("My Canvas Text tutorial")
	w.Resize(fyne.NewSize(400, 400))

	// canvaCircle := canvas.NewCircle(color.Black)
	canvaCircle := canvas.NewCircle(color.NRGBA{R: 255, G: 0, B: 255, A: 255})
	w.SetContent(canvaCircle)
	w.ShowAndRun()
}
