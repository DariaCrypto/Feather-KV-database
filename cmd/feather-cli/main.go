package main

import (
	"errors"
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

	term, err := cli.NewTerminal("Start terminal")
	if err != nil {
		log.Panicf("can not create terminal: v%", err)
	}
	term.StartSession(c)
}
