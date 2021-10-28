package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	canvasApp := app.New()

	w := canvasApp.NewWindow("My Canvas rectangle")
	rect := canvas.NewRectangle(color.NRGBA{R: 255, G: 0, B: 255, A: 255})
	rect.FillColor = color.NRGBA{R: 220, G: 255, B: 89, A: 255}
	rect.StrokeColor = color.NRGBA{R: 0, G: 255, B: 255, A: 255}
	rect.StrokeWidth = 4
	w.Resize(fyne.NewSize(400, 400))
	w.SetContent(rect)
	w.ShowAndRun()
}
