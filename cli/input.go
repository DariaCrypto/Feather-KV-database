package cli

import (
	"errors"
	"regexp"
	"strings"
)

func parseCommand(t *Terminal) (string, []string, error) {
	comm, err := t.term.ReadLine()
	if err != nil {
		return "", nil, errors.New("cli: error on command parsing: " + err.Error())
	}
	re := regexp.MustCompile(`\b\w+\b`)
	words := re.FindAllString(comm, -1)
	arguments := make([]string, 0)
	for _, str := range words {
		arguments = append(arguments, strings.TrimSpace(str))
	}
	return strings.ToUpper(strings.TrimSpace(comm)), arguments, err
}