package internal_test

import (
	"fmt"
	"json2go/internal"
	"testing"
)

func TestLoadConfig(test *testing.T) {
	var conf internal.Config
	fmt.Println(conf)
	conf.SetDefaults()

	fmt.Println(conf)
}