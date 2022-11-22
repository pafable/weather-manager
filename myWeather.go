package main

import (
	"flag"
	"fmt"
	"os"
)

type xyz struct {
	w *string
	x *string
}

var MainCmd = flag.NewFlagSet("x", flag.ExitOnError)

func main() {
	x := xyz{
		MainCmd.String("y", "", "enter a str"),
		MainCmd.String("x", "", "enter a str"),
	}

	//fmt.Println(len(os.Args))
	fmt.Printf("%s\n", os.Args[2:])
	err := MainCmd.Parse(os.Args[2:])
	fmt.Println(MainCmd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*x.w)
	if *x.x != "" {
		fmt.Println(*x.x)
	}
}
