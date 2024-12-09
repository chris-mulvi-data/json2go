/*
Loading the Configuration file and setting default values for options not defined
in the config.
*/
package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DestinationPath string `json:"destinationPath"`
	PreabmleText    string `json:"preambleText"`
}

func (c *Config) SetDefaults() {
	if c.DestinationPath == "" {
		userHome := os.Getenv("HOME")
		c.DestinationPath = userHome + "/json2go-gen"
	}
}

// LoadConfig loads the configuration file.  The expected location for the config
// is "~/.config/json2go/config.json"
func LoadConfig(conf *Config) {

	userHome := os.Getenv("HOME")
	f, err := os.ReadFile(userHome + "/.config/json2go/config.json")
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(f, &conf)
}
