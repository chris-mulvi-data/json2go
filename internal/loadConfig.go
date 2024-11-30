/*
Loading the Configuration file and setting default vaules for options not defined
in the config.
*/
package internal

type Config struct {
	DestinationPath string `json:"DestinationPath"`
}

func LoadConfig(conf *Config) {
	conf.DestinationPath = "test"
}
