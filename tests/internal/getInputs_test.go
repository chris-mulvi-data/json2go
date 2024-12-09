package internal_test

import (
	"json2go/internal"
	"os"
	// "os"
	"testing"
)

func Test_GetUserInputs(t *testing.T) {

	var sourcePath string = "test"
	var destinationPath string = ""
	os.Args = []string{sourcePath, destinationPath}

	source, destination := internal.GetUserInputs()

	if source != sourcePath {
		t.Errorf("source path read incorrectly")
	}
	if destination != destinationPath {
		t.Errorf("destination path read incorrectly")
	}
}
