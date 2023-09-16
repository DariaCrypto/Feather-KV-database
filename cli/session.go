package cli

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ddonskaya/feather/client"
	"golang.org/x/crypto/ssh/terminal"
)

type Terminal struct {
	term    *terminal.Terminal
	state   *terminal.State
	helpMsg string
}

func NewTerminal(msg string) (*Terminal, error) {
	screen := struct {
		io.Reader
		io.Writer
	}{os.Stdin, os.Stdout}

	term := terminal.NewTerminal(screen, "")
	term.SetPrompt("feather-cli >>> ")
	state, err := terminal.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return nil, fmt.Errorf("error on cli session startup: %v", err)
	}
	return &Terminal{term: term, state: state, helpMsg: msg}, nil
}

func (t *Terminal) StartSession(c *client.FeatherClient) {
	for {
		cmd, args, err := parseCommand(t)
		if err != nil {
			log.Printf("cli: can not read parse command from terminal: %v", err)
			continue
		}

		response, err := request(c, cmd, args)
		if err != nil {
			log.Printf("cli: can not request a server: %v", err)
			continue
		}
	}
}

func request(c *client.FeatherClient, cmd string, args []string) (response string, err error) {
	switch cmd {
	case "PING":
		response, err = ping(c, args)
	}

	return
}

func (t *Terminal) CloseSession() {
	terminal.Restore(0, t.state)

	log.Print("Cli session was successfully closed.")
	os.Exit(0)
}
