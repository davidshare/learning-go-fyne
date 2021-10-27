package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	canvasApp := app.New()

	w := canvasApp.NewWindow("My Canvas Text tutorial")
	w.Resize(fyne.NewSize(400, 400))
	img := canvas.NewImageFromFile("./47.jpg")
	w.SetContent(img)
	w.ShowAndRun()
}
