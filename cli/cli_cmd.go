package cli

import (
	"errors"

	"github.com/ddonskaya/feather/client"
)

func ping(c *client.FeatherClient, args []string) (string, error) {
	if len(args) != 0 {
		return "", errors.New("cli: args for PING cmd is not needed")
	}

	responce, err := 
	return responce, err
}
