package main

import (
	"flag"
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

var maxUsr = flag.Int("mu", 10, "max number of users which can connect to the chatroom at the same time")
var maxMsg = flag.Int("mm", 200, "max number of messages that will be saved in the chatroom at any time")
var maxCh = flag.Int("mc", 1024, "max number of characters for a single message")

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 1 {
		usg.Base()
		os.Exit(0)
	}

	mode := args[0]
	if mode == hlpStr {
		if len(args) == 1 {
			usg.Help()
			os.Exit(0)
		}

		str := args[1]
		if str == srvStr {
			usg.Srv()
		} else if str == clcStr {
			usg.Clc()
		} else {
			usg.UnkHelp(str)
		}
	} else if mode == srvStr {
		port := args[len(args)-1]
		pNum, err := strconv.Atoi(port)
		if err != nil || pNum <= 0 || pNum >= 65536 {
			usg.InvPort(port)
			os.Exit(1)
		}

		server := srv.BblSrv{Port: pNum, MaxUsr: *maxUsr, MaxMsg: *maxMsg, MaxCh: *maxCh}
		server.Run()
	} else if mode == clcStr {

	} else {
		usg.UnkMode(mode)
	}
}
