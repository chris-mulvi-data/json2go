package internal_test

import (
	"fmt"
	"json2go/internal"
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {

	home := os.Getenv("HOME")
	tests := []struct {
		path   string
		exists bool
	}{
		{home + "/.config", true},
		{home + "/.config/nothing", false},
		{home + "/.zshrc", true},
		{home + "/badFileName.txt", false},
	}

	for _, test := range tests {
		fileExists := internal.PathExists(test.path)
		if fileExists != test.exists {
			t.Errorf("PathExists for '%s' returned %t -> expected %t", test.path, fileExists, test.exists)
		} else {
			fmt.Printf("Expected: %t -> Received: %t\n", test.exists, fileExists)
		}
	}
}
