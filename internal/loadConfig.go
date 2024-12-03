/*
Loading the Configuration file and setting default values for options not defined
in the config.
*/
package internal

import "os"

type Config struct {
	DestinationPath string `json:"DestinationPath"`
}

func (c *Config) SetDefaults() {
	if c.DestinationPath == "" {
		userHome := os.Getenv("HOME")
		c.DestinationPath = userHome + "/json2go-gen"
	}
}

func LoadConfig(conf *Config) {
	
	conf.DestinationPath = "test"
}
