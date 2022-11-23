package weather

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"weather-manager/pkg/wmConstants"
)

var MainCmd *flag.FlagSet = flag.NewFlagSet("condition", flag.ExitOnError)

func parseWeatherArgs() (*string, error) {
	x := MainCmd.String("location", "", "enter zip code")
	err := MainCmd.Parse(os.Args[2:])
	return x, err
}

func GetCondition() *wmConstants.Wstruct {
	zip, err := parseWeatherArgs()
	if err != nil {
		log.Fatal(err)
	}

	body := wmConstants.GetCurrent(zip)

	w := wmConstants.Wstruct{}
	err = json.Unmarshal(body, &w)
	if err != nil {
		log.Fatal(err)
	}

	return &w
}
