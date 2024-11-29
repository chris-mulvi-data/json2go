package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserInputs() (source string, destination string) {
	source = ""
	destination = ""

	cliArgs := os.Args

	if len(cliArgs) < 2 {
		source = StringInput("JSON file path:")
		destination = StringInput("Destination:")
	} else {
		source = cliArgs[0]
		destination = cliArgs[1]
	}

	return source, destination
}

/*
Prompts the user for a string input.

Params:

	prompt - The string to show the user describing what is expected of them.

Returns:

	The text the user enters.

Usage:

	StringInput("The prompt to show")
*/
func StringInput(prompt string) string {

	var inputString string

	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, prompt+" ")
		inputString, _ = r.ReadString('\n')
		if inputString != "" {
			break
		}
	}
	return strings.TrimSpace(inputString)
}
