package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type TrafficStatus struct {
	on      bool
	counter uint8
}

var trafficStatus = TrafficStatus{on: true, counter: 1}

func main() {

	trafficApp := app.New()

	w := trafficApp.NewWindow("Traffic light window")
	w.Resize(fyne.NewSize(400, 400))
	fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())

	rect1 := canvas.NewCircle(color.NRGBA{R: 255, G: 255, B: 255, A: 255})
	rect1.Resize(fyne.NewSize(50, 50))
	rect2 := canvas.NewCircle(color.NRGBA{R: 255, G: 255, B: 255, A: 255})
	rect2.Resize(fyne.NewSize(50, 50))
	rect3 := canvas.NewCircle(color.NRGBA{R: 255, G: 255, B: 255, A: 255})
	rect3.Resize(fyne.NewSize(50, 50))

	btnR := widget.NewButton("Red", func() {
		goRed(rect1, rect2, rect3)
	})

	btnG := widget.NewButton("Green", func() {
		goGreen(rect1, rect2, rect3)
	})

	btnB := widget.NewButton("Blue", func() {
		goBlue(rect1, rect2, rect3)
	})

	btnStart := widget.NewButton("Start", func() {
		runTrafficLight(30, rect1, rect2, rect3)
	})

	btnStop := widget.NewButton("Stop", func() {
		trafficStatus.on = false
		Reset(rect1, rect2, rect3)
	})

	btnReset := widget.NewButton("Reset", func() {
		Reset(rect1, rect2, rect3)

		rect1.Refresh()
		rect2.Refresh()
		rect3.Refresh()
	})

	leftContainer := container.NewGridWithRows(
		7,
		layout.NewSpacer(),
		rect1,
		layout.NewSpacer(),
		rect2,
		layout.NewSpacer(),
		rect3,
	)

	rightContainer := container.NewGridWithRows(
		12,
		layout.NewSpacer(),
		btnR,
		layout.NewSpacer(),
		btnG,
		layout.NewSpacer(),
		btnB,
		layout.NewSpacer(),
		btnStart,
		layout.NewSpacer(),
		btnStop,
		layout.NewSpacer(),
		btnReset,
		layout.NewSpacer(),
	)

	w.SetContent(
		container.NewHSplit(leftContainer, rightContainer),
	)
	w.ShowAndRun()
}

func incrementMove(x uint8) (a uint8) {
	if x == 255 {
		x = 0
	} else if x == 0 {
		x = 255
	}
	return x
}

func Reset(rect1, rect2, rect3 *canvas.Circle) {
	rect1.FillColor = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	rect2.FillColor = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	rect3.FillColor = color.NRGBA{R: 255, G: 255, B: 255, A: 255}

	rect1.Refresh()
	rect2.Refresh()
	rect3.Refresh()
}

func goRed(rect1, rect2, rect3 *canvas.Circle) {
	var x uint8 = 0

	Reset(rect1, rect2, rect3)
	rect1.FillColor = color.NRGBA{R: incrementMove(x), G: 0, B: 0, A: 255}
	rect1.Refresh()
}

func goGreen(rect1, rect2, rect3 *canvas.Circle) {
	var x uint8 = 0
	Reset(rect1, rect2, rect3)
	rect2.FillColor = color.NRGBA{R: 0, G: incrementMove(x), B: 0, A: 255}
	rect2.Refresh()
}

func goBlue(rect1, rect2, rect3 *canvas.Circle) {
	var x uint8 = 0

	Reset(rect1, rect2, rect3)
	rect3.FillColor = color.NRGBA{R: 0, G: 0, B: incrementMove(x), A: 255}
	rect3.Refresh()
}

func runTrafficLight(n time.Duration, rect1, rect2, rect3 *canvas.Circle) {
	trafficStatus.on = true
	if trafficStatus.on {
		for range time.Tick(n * time.Second) {
			startTrafficLights(rect1, rect2, rect3)
			trafficStatus.counter++
		}
	}
}

func startTrafficLights(rect1, rect2, rect3 *canvas.Circle) {
	if trafficStatus.counter == 1 {
		goGreen(rect1, rect2, rect3)
	} else if trafficStatus.counter == 2 {
		goRed(rect1, rect2, rect3)
	} else if trafficStatus.counter == 3 {
		goBlue(rect1, rect2, rect3)
	} else if trafficStatus.counter > 3 {
		trafficStatus.counter = 0
	}
}
