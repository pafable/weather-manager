package weather

import (
	"flag"
	"fmt"
	"os"
)

const CURRENTWEATHER = "http://api.weatherapi.com/v1/current.json"

var APIKEY string = os.Getenv("APIKEY")
var LOCATION string = os.Getenv("LOCATION")
var MainCmd *flag.FlagSet = flag.NewFlagSet("weather", flag.ExitOnError)

type Wstruct struct {
	Current struct {
		TempC     float64 `json:"temp_c"`
		IsDay     int     `json:"is_day"`
		Condition struct {
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

func Weather() *string {
	fmt.Println(os.Args[1:])
	x := MainCmd.String("myarg", "lol", "enter zip code")
	flag.Parse()
	return x
}
