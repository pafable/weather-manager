package temperature

import (
	"flag"
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
