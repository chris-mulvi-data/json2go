package internal_test

import (
	"fmt"
	"json2go/internal"
	"os"

	// "os"
	"testing"
)

func Test_GetUserInputs(t *testing.T) {

	var sourcePath string = "test"
	var destinationPath string = ""
	os.Args = []string{sourcePath, destinationPath}

	var conf internal.Config
	conf.SetDefaults()
	fmt.Println(conf)
	source, destination := internal.GetUserInputs(&conf)

	if source != sourcePath {
		t.Errorf("source path read incorrectly")
	}
	if destination != destinationPath {
		t.Errorf("destination path read incorrectly")
	}
}
