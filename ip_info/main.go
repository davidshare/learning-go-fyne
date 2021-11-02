package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type IP struct {
	Query   string
	Country string
	City    string
}

func main() {

	ipInfo := app.New()

	w := ipInfo.NewWindow("IP info")
	w.Resize(fyne.NewSize(400, 400))
	labelTitle := widget.NewLabel("IP info")
	labelValue := widget.NewLabel("...")
	labelCountry := widget.NewLabel("...")
	labelCity := widget.NewLabel("...")

	btn := widget.NewButton("Run", func() {
		labelValue.Text = fmt.Sprintf("IP address: %s", myIP(""))
		labelValue.Refresh()

		labelCountry.Text = fmt.Sprintf("Country: %s", myIP("country"))
		labelCountry.Refresh()

		labelCity.Text = fmt.Sprintf("City: %s", myIP("city"))
		labelCity.Refresh()
	})

	w.SetContent(container.NewVBox(
		labelTitle,
		labelValue,
		labelCountry,
		labelCity,
		btn,
	))

	w.ShowAndRun()
}

func myIP(query_param string) string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}

	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}
	var ip IP
	json.Unmarshal(body, &ip)

	switch param := query_param; param {
	case "country":
		return ip.Country

	case "city":
		return ip.City

	default:
		return ip.Query
	}
}
