package main

import (
	"fmt"
	"os"

	"github.com/hitenpratap/go-web-server/servers"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a server type: nethttp | gin | echo | fiber")
		return
	}

	switch os.Args[1] {
	case "nethttp":
		servers.StartNetHTTPServer()
	case "gin":
		servers.StartGinServer()
	case "echo":
		servers.StartEchoServer()
	case "fiber":
		servers.StartFiberServer()
	default:
		fmt.Println("Unknown server type. Please use: nethttp | gin | echo | fiber")
	}
}
