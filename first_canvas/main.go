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
	colorX := color.NRGBA{R: 0, G: 255, B: 0, A: 255}
	textX := canvas.NewText("This is the text", colorX)
	w.SetContent(textX)
	w.ShowAndRun()
}
