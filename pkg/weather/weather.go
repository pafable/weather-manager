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
	var randNum int = rand.Intn(100)

	switch {
	case randNum < 33:
		return fmt.Sprintf("The winds are slow and are blowing at %d mph", randNum)
	case randNum < 66:
		return fmt.Sprintf("The winds are med and are blowing at %d mph", randNum)
	}
	return fmt.Sprintf("The winds are fast and are blowing at %d mph", randNum)
}
