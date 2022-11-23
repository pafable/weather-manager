package errChkr

import (
	"log"
	"os"
)

func ErrorCheck() {
	if len(os.Args) < 2 {
		log.Fatal("condition or temperature is required!")
	}

	if os.Args[1] == "temperature" || os.Args[1] == "condition" {
	} else {
		log.Fatal("condition or temperature is required!")
	}
}
