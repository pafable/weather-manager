package temperature

import (
	"flag"
)

type Temp struct {
	Temperature *int
}

var MainCmd = flag.NewFlagSet("temp", flag.ExitOnError)

func Temperature() *int {
	return MainCmd.Int("degree", 0, "enter the temperature in celsius")
}
