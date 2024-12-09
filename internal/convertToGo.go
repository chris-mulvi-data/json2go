/*
Convert To Go will take the json file that has been read
from the user and turn them into a struct.

The output will be written to a new file.
*/

package internal

import (
	"os"
)

/*
Handle error checking
*/
func check(e error) {
	if e != nil {
		panic(e)
	}
}

/*
WriteGoFile creates a new .go file that the struct(s) will be created into
based on the JSON file contents.

Example usage:

	WriteGoFile("path/to/file.go", {"some\n", "lines\n"})
*/
func WriteGoFile(destinationPath string, content []string) {

	file, err := os.Create(destinationPath)
	check(err)

	// Ensure clean close when done.
	defer file.Close()

	//write the preamble to the file.
	preamble := getPreable()
	file.WriteString(preamble)
	file.WriteString("package main\n\n")

	// write the lines into the file
	for _, line := range content {
		file.WriteString(line)
	}
}

/*
GetPreable get a string of text that is used as a file level GoDoc describing
that the contents were generated from JSON
*/
func getPreable() string {
	var config Config
	LoadConfig(&config)
	return config.PreabmleText
}
