package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fmt.Println("Test Fyne...")

	a := app.New()
	w := a.NewWindow("Welcome to fyne")

	hello := widget.NewLabel("Hello fyne")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :")
		}),
	))
	w.ShowAndRun()
}
