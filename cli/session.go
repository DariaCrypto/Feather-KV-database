package cli

import (
	"log"
	"os"

	"github.com/ddonskaya/feather/client"

	"golang.org/x/crypto/ssh/terminal"
)

type Terminal struct {
	state   *terminal.State
	helpMsg string
}

func NewTerminal(msg string) *Terminal {
	state, err := terminal.MakeRaw(0)
	if err != nil {
		log.Fatalf("Error on cli session startup: %s. Exiting.", err.Error())
	}

	return &Terminal{state: state, helpMsg: msg}
}

func (t *Terminal) StartSession(c *client.FeatherClient) {
	for {
		cmd, args, err := parseCommand()
		if err != nil {
			continue
		}

		request(c, cmd, args)
		if err != nil {
			continue
		}
	}
}

func request(c *client.FeatherClient, cmd string, args []string) (string, error) {
	switch cmd {
	case "PING":
		ping(c, args)
	}

	return "", nil
}

func (t *Terminal) CloseSession() {
	terminal.Restore(0, t.state)

	log.Print("Cli session was successfully closed.")
	os.Exit(0)
}
