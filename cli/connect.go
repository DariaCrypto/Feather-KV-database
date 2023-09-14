package cli

import (
	"github.com/ddonskaya/feather/client"
	"fmt"
	"log"
)

func ObtainClient(host string, port int) (*client.FeatherClient, error) {
	addr := fmt.Sprintf("%s:%d", host, port)
	return connect(addr, "tcp")
}

func connect(addr, network string) (*client.FeatherClient, error) {
	c := client.NewFeatherClient()

	if _, err := client.Ping(c); err != nil {
		log.Printf("cli: connection error: %v", err)
		return nil, err
	}

	log.Printf("cli: connection is successful. Address connection is %s", addr)
	return c, nil
}
