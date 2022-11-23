package main

import (
	"fmt"
	"os"
	"strings"
	"weather-manager/pkg/errChkr"
	"weather-manager/pkg/temperature"
	"weather-manager/pkg/weather"
)

func main() {
	//  checks to make sure there are args
	errChkr.ErrorCheck()

	switch os.Args[1] {
	case "condition":
		w := weather.GetCondition()
		fmt.Printf("It is %s with winds blowing at %.1f mph from the %s\n", strings.ToLower(w.Current.Condition.Text), w.Current.WindMph, w.Current.WindDir)

	case "temperature":
		t := temperature.GetTemp()
		fmt.Printf("It is %.1f C, but feels like %.1f C\n", t.Current.TempC, t.Current.FeelslikeC)
		fmt.Printf("It is %.1f F, but feels like %.1f F\n", temperature.GetFahrenheit(t.Current.TempC), t.Current.FeelslikeF)
		fmt.Printf("It is %.1f K\n", temperature.GetKelvin(t.Current.TempC))
	}
}
