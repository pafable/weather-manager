package condition

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"weather-manager/pkg/wmConstants"
)

var MainCmd *flag.FlagSet = flag.NewFlagSet("condition", flag.ExitOnError)

func parseWeatherArgs() *string {
	x := MainCmd.String("location", "", "enter zip code")
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

func GetCondition() *wmConstants.Wstruct {
	zip := parseWeatherArgs()

	body := wmConstants.GetCurrent(zip)

	w := wmConstants.Wstruct{}
	err := json.Unmarshal(body, &w)
	if err != nil {
		log.Fatal(err)
	}

	return &w
}
