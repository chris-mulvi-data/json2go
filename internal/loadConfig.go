/*
Loading the Configuration file and setting default values for options not defined
in the config.
*/
package internal

import (
	"fmt"
	"os"
)

type Config struct {
	DestinationPath string `json:"destinationPath"`
}

func (c *Config) SetDefaults() {
	if c.DestinationPath == "" {
		userHome := os.Getenv("HOME")
		c.DestinationPath = userHome + "/json2go-gen"
	}
}

func LoadConfig(conf *Config) {

	f, err := os.Open("")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
}
