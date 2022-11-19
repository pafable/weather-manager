package errChkr

import (
	"log"
	"os"
)

func ErrorCheck() {
	if len(os.Args) < 2 {
		log.Fatal("weather or temperature is required!")
	}

	if os.Args[1] == "temp" || os.Args[1] == "weather" {
	} else {
		log.Fatal("weather or temperature is required!")
	}
}
