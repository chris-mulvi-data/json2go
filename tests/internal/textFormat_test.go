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

func TestPadRightToSize(t *testing.T) {
	tests := []string{"test", "tested", "testing"}

	for _, test := range tests {
		i.PadRightToSize(&test, 7)
		fmt.Printf("%s lined up\n", test)
	}
}

func TestAlignColumns(t *testing.T) {
	input := []string{
		"Bears like to eat",
		"kake and hunny",
		"with some nesquick",
	}

	expected := []string{
		"Bears like to       eat",
		"kake  and  hunny",
		"with  some nesquick",
	}

	fmt.Println("starting:")
	for i := range input {
		fmt.Println(input[i])
	}
	fmt.Println()

	results := i.AlignColumns(input)

	fmt.Println("result:")
	for i := 0; i < len(input); i++ {
		if results[i] != expected[i] {
			t.Errorf("expected: \"%s\", got: \"%s\"\n", expected[i], results[i])
		}
		fmt.Println(results[i])
	}
}
