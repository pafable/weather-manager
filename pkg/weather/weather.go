package weather

import (
	"flag"
)

type Wthr struct {
	Weather *string
}

var MainCmd *flag.FlagSet = flag.NewFlagSet("weather", flag.ExitOnError)

func Weather() *string {
	return MainCmd.String("cond", "", "enter the condition")
}
