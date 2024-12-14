package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Gets the source and destination file paths.  The source path is the JSON file
to be used to create the Go struct that will get stored at the destination file path.

If no CLI arguments are supplied, the user is prompted to provide the source path and
the destination path.  If only one CLI argument is supplied, it is assumed to be the
source json file and a default path will be used as the destination.

When prompted, the user will also have the option to use a default destination.
*/
func GetUserInputs(conf *Config) (source string, destination string) {
	source = ""
	destination = ""

	cliArgs := os.Args

	switch len(cliArgs) {
	case 0:
		// prompt for source and destination
		source = StringInput("JSON file path:")
		destination = StringInput("Destination:")

	case 1:
		// use the only arg as the source file
		// use a default destination file path
		source = cliArgs[0]

	case 2:
		// use the first arg as the source file
		// use the second arg as the destination file path
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
