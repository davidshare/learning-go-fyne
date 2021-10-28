package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("Gradient : canvas")
	w.Resize(fyne.NewSize(400, 400))
	// gradient1 := canvas.NewHorizontalGradient(color.White, color.Black)
	// gradient2 := canvas.NewLinearGradient(color.White, color.Black, 45)
	// gradient3 := canvas.NewRadialGradient(color.White, color.Black)
	gradient4 := canvas.NewVerticalGradient(color.White, color.Black)

	w.SetContent(gradient4)

	w.ShowAndRun()
}
