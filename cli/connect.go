package cli

import (
	"feather-kv/client"
	"fmt"
	"log"

)

func ObtainClient(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)
	return connect(addr, "tcp")
}

func connect(addr, network string) (*client.FeatherClient, error) {
	client := client.New()

	if _, err := c.Ping() // check err connection
	
	log.Println("cli: connection is successful. Address connection is ", addr)
	return client, nil
}