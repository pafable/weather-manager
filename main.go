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

		fmt.Printf("The weather condition is %s\n", *w1.Weather)
		fmt.Printf("%s\n", w1.WindCond())

	case "temp":
		// checks to make sure args are passed into temp flag
		errChkr.TempParamCheck()

		fmt.Printf("The temperature is %.1f celsius\n", *t1.Temperature)

		if *t1.Fahrenheit() >= 88.0 {
			fmt.Printf("The temperature is %.1f fahrenheit, it's burning hot!\n", *t1.Fahrenheit())
		} else {
			fmt.Printf("The temperature is %.1f fahrenheit\n", *t1.Fahrenheit())
		}

		fmt.Printf("The temperature is %.1f kelvin\n", *t1.Kelvin())
	}
}
