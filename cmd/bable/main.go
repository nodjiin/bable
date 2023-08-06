package main

import (
	"os"

	"github.com/nodjiin/bable/internal/pkg/usg"
)

const (
	hlpStr = "help"
	srvStr = "server"
	clcStr = "client"
)

func main() {
	if len(os.Args) == 1 {
		usg.Base()
		os.Exit(0)
	}

	mode := os.Args[1]
	if mode == hlpStr {
		if len(os.Args) == 2 {
			usg.Help()
			os.Exit(0)
		}

		str := os.Args[2]
		if mode == srvStr {
			usg.Srv()
		} else if mode == clcStr {
			usg.Clc()
		} else {
			usg.UnkHelp(str)
		}
	} else if mode == srvStr {

	} else if mode == clcStr {

	} else {
		usg.UnkMode(mode)
	}
}
