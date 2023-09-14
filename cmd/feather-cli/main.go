package main

import (
	"log"

	"github.com/ddonskaya/feather/cli"
	"github.com/ddonskaya/feather/client"
)

func main() {

	var (
		c   *client.FeatherClient
		err error
	)

	host := "127.0.0.1"
	port := 6870

	c, err = cli.ObtainClient(host, port)

	if err != nil {
		log.Printf("cli: %v", err)
		return
	}

	term := cli.NewTerminal("Start terminal")
	term.StartSession(c)
}
