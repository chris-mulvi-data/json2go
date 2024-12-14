package main

import (
	"fmt"
	"json2go/internal"
	"os"
)

func main() {

	/*
		todo: check for and load settings found in a "~/.config/json2go/config.json" file
		 	  this includes overrides for the default destination file path.
	*/
	var conf internal.Config

	userHome := os.Getenv("HOME")
	var configExists bool = internal.PathExists(userHome + "/.config/json2go/config.json")

	if configExists {
		internal.LoadConfig(&conf)
	}
	// set defaults for any config values that were not found in
	// the user's config file.
	conf.SetDefaults()

	source, destination := internal.GetUserInputs(&conf)

	fmt.Println("The source is:", source)
	fmt.Println("The destination is:", destination)
}
