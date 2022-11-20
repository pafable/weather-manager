package main

import (
	"fmt"
	"os"
	"weather-manager/pkg/errChkr"
	"weather-manager/pkg/temperature"
	"weather-manager/pkg/weather"
)

func main() {
	//  checks to make sure there are args
	errChkr.ErrorCheck()

	w1 := weather.Wthr{
		Weather: weather.Weather(),
	}

	t1 := temperature.Temp{
		Temperature: temperature.Temperature(),
	}

	switch os.Args[1] {
	case "weather":
		// checks to make sure args are passed into weather flag
		errChkr.WeatherParamCheck()

		w1.Result()

		resWthr, resWind := w1.Result()
		fmt.Printf("%s\n", *resWthr)
		fmt.Printf("%s\n", resWind)

	case "temp":
		// checks to make sure args are passed into temp flag
		errChkr.TempParamCheck()

		celsius, fahrenheit, kelvin := t1.Result()
		fmt.Printf("%s", *celsius)
		fmt.Printf("%s", *fahrenheit)
		fmt.Printf("%s", *kelvin)
	}
}
