package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"

	"encoding/json"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type IP struct {
	Query   string
	Country string
	City    string
}

type Welcome struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int64     `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int64     `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int64     `json:"timezone"`
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Cod        int64     `json:"cod"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
}

type Sys struct {
	Type    int64  `json:"type"`
	ID      int64  `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type Weather struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed int64 `json:"speed"`
	Deg   int64 `json:"deg"`
}

func main() {

	weatherApp := app.New()

	w := weatherApp.NewWindow("Weather api & fyne")
	w.Resize(fyne.NewSize(400, 400))

	res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=peshawar&APPID=88a3325d8b543b9103c71abe0ebc15ef")

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	weather, err := UnmarshalWelcome(body)

	if err != nil {
		fmt.Println(err)
	}

	appTitle := canvas.NewText("Weather API && fyne", color.Black)
	appTitle.TextStyle = fyne.TextStyle{Bold: true}
	country := canvas.NewText(fmt.Sprintf("Country: %s", weather.Sys.Country), color.Black)
	wind := canvas.NewText(fmt.Sprintf("wind speed: %d", weather.Wind.Speed), color.Black)
	temp := canvas.NewText(fmt.Sprintf("TEMP: %.2f", weather.Main.Temp), color.Black)

	w.SetContent(container.NewVBox(
		appTitle,
		country,
		wind,
		temp,
	))

	w.ShowAndRun()
}

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    welcome, err := UnmarshalWelcome(bytes)
//    bytes, err = welcome.Marshal()

func UnmarshalWelcome(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
