package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

func main() {

	dice := app.New()

	w := dice.NewWindow("New card")
	w.Resize(fyne.NewSize(400, 400))

	img := canvas.NewImageFromFile("./dice/dice6.png")

	w.SetContent(widget.NewCard(
		"Card title",
		"Cart subtitme",
		img,
	))

	w.ShowAndRun()
}
