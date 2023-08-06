package main

import (
	"os"
	"strconv"

	"github.com/nodjiin/bable/internal/pkg/srv"
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
		// TODO flags
		var server srv.BblSrv
		port := os.Args[2] // change to account flags

		pNum, err := strconv.Atoi(port)
		if err == nil && (pNum >= 0 && pNum <= 65536) {
			usg.InvPort(port)
			os.Exit(1)
		}

		server.Port = pNum
		server.Run()
	} else if mode == clcStr {

	} else {
		usg.UnkMode(mode)
	}
}
