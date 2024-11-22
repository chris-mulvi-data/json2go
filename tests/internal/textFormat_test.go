package internal_test

import (
	"fmt"
	i "json2go/internal"
	"testing"
)

func TestPropertyToField(t *testing.T) {
	tests := []struct {
		input     string
		expected  string
		wantError bool
	}{
		{"json_property_name", "JsonPropertyName", false},
		{"jsonpropertyname", "JsonPropertyName", true},
	}

	for _, test := range tests {
		result, err := i.PropertyToField(test.input)
		if err != nil && test.wantError == false {
			t.Errorf("PropertyToField() error: %v", err)
		} else if result != test.expected && test.wantError == false {
			t.Errorf("PropertyToField() -> expected: %v, got: %v", test.expected, result)
		}
		fmt.Printf("Expected: %v\nGot: %v\n\n", test.expected, result)
	}

}
