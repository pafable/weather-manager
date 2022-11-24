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

	if len(os.Args) < 3 {
		log.Fatal("needs more args (see readme.md)")
	}

	if (os.Args[1] == "temperature" || os.Args[1] == "condition") && (os.Args[2] == "temperature" || os.Args[2] == "condition") {
		log.Fatal("choose only one, condition or temperature")
	}
}
