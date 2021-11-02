package main

import (
	"fmt"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	dice := app.New()

	w := dice.NewWindow("Dice game window")
	w.Resize(fyne.NewSize(400, 400))

	img := canvas.NewImageFromFile("./dice/dice6.png")
	img.FillMode = canvas.ImageFillOriginal

	btn1 := widget.NewButton("Play", func() {
		rand := rand.Intn(6) + 1
		img.File = fmt.Sprintf("./dice/dice%d.png", rand)
		img.Refresh()
	})

	w.SetContent(container.NewVBox(
		img,
		btn1,
	))

	w.ShowAndRun()
}
