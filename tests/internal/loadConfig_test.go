package internal_test

import (
	"json2go/internal"
	"log"
	"testing"
)

func TestLoadConfig(test *testing.T) {
	log.Println("Starting TestLoadConfig")
	var conf internal.Config
	log.Printf("Config object before loading:\n%v\n\n", conf)
	internal.LoadConfig(&conf)
	log.Printf("Config object after loading:\n%v\n\n", conf)
	conf.SetDefaults()
	log.Printf("Config object after loading defaults:\n%v\n\n", conf)
}
