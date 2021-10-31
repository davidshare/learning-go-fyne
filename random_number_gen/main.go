package main

import (
	"fmt"
	"image/color"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Random number generator")
	w.Resize(fyne.NewSize(400, 400))

	label1 := canvas.NewText("Rand number gen", color.Black)
	label1.TextSize = 40

	btn1 := widget.NewButton("Generate", func() {
		rand1 := rand.Intn(1000000)
		label1.Text = fmt.Sprint(rand1)
		label1.Refresh()
	})

	w.SetContent(
		container.NewVBox(
			label1,
			btn1,
		),
	)
	w.ShowAndRun()

}
