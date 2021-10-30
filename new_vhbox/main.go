package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()

	w := a.NewWindow("Layout NewVBox and NewHBox")
	w.Resize(fyne.NewSize(400, 400))

	btn1 := widget.NewButton("Click me", func() {

	})

	label1 := widget.NewLabel("here is my text")

	// box1 := container.NewVBox(
	// 	btn1,
	// 	label1,
	// )

	box2 := container.NewHBox(btn1, label1)

	w.SetContent(box2)

	w.ShowAndRun()
}
