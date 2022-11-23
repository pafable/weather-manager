package temperature

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"weather-manager/pkg/wmConstants"
)

var MainCmd = flag.NewFlagSet("temperature", flag.ExitOnError)

func parseTempArgs() (*string, error) {
	x := MainCmd.String("location", "0", "enter zip code")
	err := MainCmd.Parse(os.Args[2:])
	return x, err

}

func GetTemp() *wmConstants.Wstruct {
	zip, err := parseTempArgs()
	if err != nil {
		log.Fatal(err)
	}

	body := wmConstants.GetCurrent(zip)

	t := wmConstants.Wstruct{}
	err = json.Unmarshal(body, &t)
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
