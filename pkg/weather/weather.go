package weather

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"weather-manager/pkg/wmConstants"
)

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

func getCurrent(s *string) []byte {
	var url string = fmt.Sprintf(
		"%s?key=%s&q=%s",
		wmConstants.CURRENTWEATHER,
		wmConstants.APIKEY,
		*s)
	resp, err := http.Get(url)

	// close connection to prevent memory leaks
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}

func Condition() *Wstruct {
	x := MainCmd.String("location", "", "enter zip code")
	err := MainCmd.Parse(os.Args[2:])
	if err != nil {
		log.Fatal(err)
	}

	body := getCurrent(x)

	w := Wstruct{}
	err = json.Unmarshal(body, &w)
	if err != nil {
		log.Fatal(err)
	}

	return &w
}
