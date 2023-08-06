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

	server		starts a local chatroom instance
	client		connect to an existing chatroom 
	
Use bable help <mode> for more information about a mode.`

	fmt.Println(str)
}

func Help() {
	Base()
}

func Srv() {
	str := `Usage: bable server [flags] <port>
	
Starts a new server instance listening to <port>.
The <port> must be open, inbound and outbound TCP traffic must be allowed.

Available flags:

	-mu		max number of users which can connect to the chatroom at the same time. Default is 10.
	-mm		max number of messages that will be saved in the chatroom at any time. Default is 200.
	-mc		max number of characters for a single message. Default is 1024.`
	fmt.Println(str)
}

func Clc() {
	str := `Usage: bable client <username>@<chatroom-ip>:<port>

Attempts to connect to an existing chatroom instance.`

	fmt.Println(str)
}

func UnkMode(mode string) {
	fmt.Printf("%s: Unknown mode\n Run 'bable' for help on usage\n", mode)
}

func UnkHelp(topic string) {
	fmt.Printf("Unknown help topic '%s'. Run 'bable help'\n", topic)
}

func InvPort(port string) {
	fmt.Printf("Invalid port value %s. Port must be an integer between 0 and 65536.\n", port)
}
