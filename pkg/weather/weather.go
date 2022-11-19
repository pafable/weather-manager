package weather

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

type Wthr struct {
	Weather *string
}

var MainCmd *flag.FlagSet = flag.NewFlagSet("weather", flag.ExitOnError)

func Weather() *string {
	return MainCmd.String("cond", "", "enter the condition")
}

func (w Wthr) WindCond() string {
	// sets random seed
	rand.Seed(time.Now().UnixNano())
	// picks a random number from a range of 100
	var randNum int = rand.Intn(100)

	switch {
	case randNum < 33:
		return fmt.Sprintf("The winds are slow and are moving at %d mph", randNum)
	case randNum < 66:
		return fmt.Sprintf("The winds are med and are breezing at %d mph", randNum)
	default:
		return fmt.Sprintf("The winds are fast and are howling at %d mph", randNum)
	}
}
