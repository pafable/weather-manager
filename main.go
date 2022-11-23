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

	t1 := temperature.Temp{
		Temperature: temperature.Temperature(),
	}

	switch os.Args[1] {
	case "condition":
		w := weather.Condition()
		fmt.Printf("It is %s with winds blowing at %.1f mph from the %s\n", strings.ToLower(w.Current.Condition.Text), w.Current.WindMph, w.Current.WindDir)
		fmt.Printf("It feels like %.1f F\n", w.Current.FeelslikeF)

	case "temp":
		celsius, fahrenheit, kelvin := t1.Result()
		fmt.Printf("%s", *celsius)
		fmt.Printf("%s", *fahrenheit)
		fmt.Printf("%s", *kelvin)
	}
}
