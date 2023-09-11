package main

import (
	"log"

	"github.com/ddonskaya/feather/cli"
	"github.com/ddonskaya/feather/client"
)

var (
	host   string
	port   int
	socket string
)

func main() {

	var (
		c *client.FeatherClient
		err error
	)
	
	if socket == "" {
		c, err = cli.ObtainClient(host, port)
	}
	if err != nil {
		log.Printf("clie: v%", err)
		return
	}

	term := cli.NewTerminal("Start terminal")
	term.StartSession(c)
}
