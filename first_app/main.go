package main

import (
	"fmt"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fmt.Println("Test Fyne...")

	a := app.New()
	w := a.NewWindow("Welcome to fyne")
	hello := widget.NewLabel("Hello fyne")
	checkBox := widget.NewCheck("Is this morning?", func(b bool) {
		if b == true {
			hello.SetText("Good morning!!!")
		} else {
			hello.SetText("Good day!!!")
		}
	})
	btn := widget.NewButton("Hi!", func() {
		hello.SetText("Welcome :")
	})

	url, _ := url.Parse("https://github.com/davidshare")
	hyperlink := widget.NewHyperlink("See my github profile", url)

	w.SetContent(container.NewVBox(
		hello,
		checkBox,
		btn,
		hyperlink,
	))
	w.Resize(fyne.NewSize(200, 400))
	w.ShowAndRun()
}
