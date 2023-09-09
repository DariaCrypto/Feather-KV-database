package cli

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/ddonskaya/feather/client"

	"golang.org/x/crypto/ssh/terminal"
)

type Terminal struct {
	state *terminal.State
	helpMsg string
}

func NewTerminal(ste *terminal.State, msg string) *Terminal {
	return &Terminal{state: ste, helpMsg: msg}
}

func StartSession(){
	state, err := terminal.MakeRaw(0)
	if err != nil {
		log.Fatalf("Error on cli session startup: %s. Exiting.", err.Error())
	}
	cli := NewTerminal(state, "")
	prepareSession()

	for {

	}
}

func Request(c *client.FeatherClient, cmd string, args []string) (string, error) {
	switch cmd{
	case "PING":
	}
}


func (t *Terminal)closeSession() {
	terminal.Restore(0, t.state)

	log.Print("Cli session was successfully closed.")
	os.Exit(0)
}