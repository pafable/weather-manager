package temperature

import (
	"flag"
	"fmt"
)

type Temp struct {
	Temperature *float64
}

var MainCmd = flag.NewFlagSet("temp", flag.ExitOnError)

func Temperature() *float64 {
	return MainCmd.Float64("degree", 0, "enter the temperature in celsius")
}

func (t Temp) Fahrenheit() *float64 {
	c := *t.Temperature
	f := (float64(c) * 1.8) + 32
	return &f
}

func (t Temp) Kelvin() *float64 {
	c := *t.Temperature
	k := float64(c) + 273.15
	return &k
}

func (t Temp) Result() (*string, *string, *string) {
	c := fmt.Sprintf("The temperature is %.1f celsius\n", *t.Temperature)
	k := fmt.Sprintf("The temperature is %.1f kelvin\n", *t.Kelvin())

	if *t.Fahrenheit() >= 88.0 {
		f := fmt.Sprintf("The temperature is %.1f fahrenheit, it's burning hot!\n", *t.Fahrenheit())
		return &c, &f, &k
	} else {
		f := fmt.Sprintf("The temperature is %.1f fahrenheit\n", *t.Fahrenheit())
		return &c, &f, &k
	}
}
