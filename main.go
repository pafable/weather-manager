package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"weather-manager/pkg/temperature"
	"weather-manager/pkg/weather"
)

var W *flag.FlagSet = weather.MainCmd
var T *flag.FlagSet = temperature.MainCmd

func errorCheck() {
	if len(os.Args) < 2 {
		log.Fatal("weather or temperature is required!")
	}

	if os.Args[1] == "temp" || os.Args[1] == "weather" {
	} else {
		log.Fatal("weather or temperature is required!")
	}
}

func main() {
	w1 := weather.Wthr{
		weather.Weather(),
	}
	t1 := temperature.Temp{
		temperature.Temperature(),
	}

	errorCheck()

	switch os.Args[1] {
	case "weather":
		if len(os.Args) < 4 {
			weather.MainCmd.PrintDefaults()
			os.Exit(1)
		}

		wErr := W.Parse(os.Args[2:])
		if wErr != nil {
			log.Fatal(wErr)
		}

		fmt.Printf("The weather condition is %s\n", *w1.Weather)

	case "temp":
		if len(os.Args) < 4 {
			temperature.MainCmd.PrintDefaults()
			os.Exit(1)
		}

		tErr := T.Parse(os.Args[2:])
		if tErr != nil {
			log.Fatal(tErr)
		}

		fmt.Printf("The temperature is %d celsius\n", *t1.Temperature)
	}
}
