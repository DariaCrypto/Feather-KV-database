package cli

import (
	"fmt"
	"regexp"
	"strings"
)

func parseCommand(t *Terminal) (string, []string, error) {
	comm, err := t.term.ReadLine()
	if err != nil {
		return "", nil, fmt.Errorf("cli: error on command parsing: %v", err)
	}
	re := regexp.MustCompile(`\b\w+\b`)
	words := re.FindAllString(comm, -1)
	arguments := make([]string, 0)
	for _, str := range words {
		arguments = append(arguments, strings.TrimSpace(str))
	}
	if len(arguments) == 1 {
		return strings.ToUpper(strings.TrimSpace(arguments[0])), []string{}, err
	}
	return strings.ToUpper(strings.TrimSpace(arguments[0])), arguments[1:], err
}