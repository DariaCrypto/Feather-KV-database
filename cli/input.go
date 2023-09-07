package cli
import (
	"os"
	"golang.org/x/crypto/ssh/terminal"
	"errors"
	"strings"
	"regexp")

var term = terminal.NewTerminal(os.Stdin, "feather>> ")


func parseCommand(commStr string) (string, []string, error) {
	comm, err := term.ReadLine()
	if err != nil {
		return "", nil, errors.New("cli: error on command parsing: " + err.Error())
	}
	re := regexp.MustCompile(`\b\w+\b`)
	words := re.FindAllString(comm, -1)
	arguments := make([]string, 0)
	for _, str := range args {
		arguments = append(arguments, strings.TrimSpace(str))
	}

	return strings.ToUpper(strings.TrimSpace(comm)), arguments, err
}