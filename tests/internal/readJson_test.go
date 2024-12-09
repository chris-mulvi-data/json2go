package internal_test

import (
	"fmt"
	"json2go/internal"
	"testing"
)

func Test_readFile(t *testing.T) {
	tests := []struct {
		filePath  string
		wantError bool
	}{
		{"../../testData/testJson.json", false},
		{"../../testData/testBadJson.json", true},
		{"nonExistent.json", true},
	}

	for _, test := range tests {
		data, err := internal.ReadFile(test.filePath)
		if err != nil && !test.wantError {
			t.Errorf("ReadFile() failed with error: %v\n", err)
		}
		if err != nil && test.wantError {
			fmt.Printf("successful failure for: %v\n", err)
		} else {
			fmt.Println(data)
		}
	}

}
