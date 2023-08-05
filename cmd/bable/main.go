package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	srvStr = "server"
	clcStr = "client"
)

func main() {
	if len(os.Args) == 1 {
		// TODO print mode/argument description
		fmt.Println("Usage 'bable <mode> [arguments]'")
		os.Exit(0)
	}

	mode := os.Args[1]
	if mode == srvStr {

	} else if mode == clcStr {

	} else {
		fmt.Printf("%s: Unknown mode\n Run 'bable' for help on usage", strings.Join(os.Args[:2], " "))
		os.Exit(1)
	}
}
