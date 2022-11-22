package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var APIKEY string = os.Getenv("APIKEY")
var LOCATION string = os.Getenv("LOCATION")

const CURRENTWEATHER = "http://api.weatherapi.com/v1/current.json"

type WeatherStruct struct {
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		Uv         float64 `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
	}
}

func main() {
	var x string = fmt.Sprintf(
		"%s?key=%s&q=%s",
		CURRENTWEATHER,
		APIKEY,
		LOCATION)
	resp, err := http.Get(x)

	// close connection to prevent memory leaks
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	w := WeatherStruct{}
	err = json.Unmarshal(body, &w)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%.1fC\n", w.Current.TempC)
	fmt.Printf("%.1fF", w.Current.TempF)
}
