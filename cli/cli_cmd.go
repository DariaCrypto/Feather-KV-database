package cli

import (
	"errors"
	"fmt"

	"github.com/ddonskaya/feather/client"
)

func ping(c *client.FeatherClient, args []string) (string, error) {
	if len(args) != 0 {
		return "", errors.New("cli: args for PING cmd is not needed")
	}

	response, err := client.Ping(c)
	if err != nil {
		return "", fmt.Errorf("cli: can not execute command PING: %v", err)
	}

	return response.Values[0], err
}
