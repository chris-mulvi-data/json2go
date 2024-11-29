package internal

import "os"

func GetUserInputs() (source string, destination string) {
	source = ""
	destination = ""

	cliArgs := os.Args

	if len(cliArgs) < 2 {
		// prompt for source
		// prompt for destination
	} else {
		source = cliArgs[0]
		destination = cliArgs[1]
	}

	return source, destination
}

// Gets the path of the source file.
func getSourcePath() string {
	return ""
}
