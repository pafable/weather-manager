package errChkr

import (
	"flag"
	"log"
	"os"
	"weather-manager/pkg/temperature"
	"weather-manager/pkg/weather"
)

func ErrorCheck() {
	if len(os.Args) < 2 {
		log.Fatal("weather or temperature is required!")
	}

	if os.Args[1] == "temp" || os.Args[1] == "weather" {
	} else {
		log.Fatal("weather or temperature is required!")
	}
}

func TempParamCheck() {
	var T *flag.FlagSet = temperature.MainCmd

	if len(os.Args) < 4 {
		T.PrintDefaults()
		os.Exit(1)
	}

	tErr := T.Parse(os.Args[2:])
	if tErr != nil {
		log.Fatal(tErr)
	}
}

func WeatherParamCheck() {
	var W *flag.FlagSet = weather.MainCmd

	if len(os.Args) < 4 {
		W.PrintDefaults()
		os.Exit(1)
	}

	wErr := W.Parse(os.Args[2:])
	if wErr != nil {
		log.Fatal(wErr)
	}
}
