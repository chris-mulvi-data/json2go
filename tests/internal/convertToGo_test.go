package internal_test

import (
	"fmt"
	"json2go/internal"
	"os"
	"testing"
)

func Test(t *testing.T) {
	t.Skip("NOOP")
}

func TestWriteGoFile(t *testing.T) {
	lines := []string{"var one string\n", "var two string\n"}
	home := os.Getenv("HOME")
	destination := fmt.Sprintf("%s/test.go", home)
	internal.WriteGoFile(destination, lines)
}
