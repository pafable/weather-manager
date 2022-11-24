package temperature

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"weather-manager/pkg/wmConstants"
)

var MainCmd = flag.NewFlagSet("temperature", flag.ExitOnError)

func parseTempArgs() *string {
	x := MainCmd.String("location", "", "enter zip code")
	MainCmd.String("scale", "celsius", "temperature scale")
	err := MainCmd.Parse(os.Args[2:])
	if err != nil {
		MainCmd.PrintDefaults()
		log.Fatal(err)
	}

	if len(os.Args) <= 2 {
		MainCmd.PrintDefaults()
		os.Exit(1)
	}

	return x
}

func GetTemp() *wmConstants.Wstruct {
	zip := parseTempArgs()
	body := wmConstants.GetCurrent(zip)

	t := wmConstants.Wstruct{}
	err := json.Unmarshal(body, &t)
	if err != nil {
		log.Fatal(err)
	}

	return &t
}

func GetFahrenheit(c float64) float64 {
	return (c * 1.8) + 32
}

func GetKelvin(c float64) float64 {
	return c + 273.15
}
