// usg package prints usage information
package usg

import (
	"fmt"
)

func Base() {
	str := `Bable is a lightweight tool to allow creation and comunication with encrypted chatrooms  

Usage:

	bable <mode> [arguments]
	
The modes are:

	server	starts a local chatroom instance
	client	connect to an existing chatroom 
	
Use bable help <mode> for more information about a mode.`

	fmt.Println(str)
}

func Help() {
	Base()
}

func Srv() {
	str := `Usage: bable server <port>
	
Starts a new server instance listening to <port>.
The <port> must be open, inbound and outbound TCP traffic must be allowed.`
	fmt.Println(str)
}

func Clc() {
	str := `Usage: bable client <username>@<chatroom-ip>:<port>

Attempts to connect to an existing chatroom instance.`

	fmt.Println(str)
}

func UnkMode(mode string) {
	fmt.Printf("%s: Unknown mode\n Run 'bable' for help on usage", mode)
}

func UnkHelp(topic string) {
	fmt.Printf("Unknown help topic '%s'. Run 'bable help'", topic)
}
