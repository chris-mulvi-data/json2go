package main

import (
	"fmt"
	"json2go/internal"
)

func main() {

	/*
		todo: check for and load settings found in a "~/.config/json2go/config.yml" file
		 	  this includes overrides for the default destination file path.
	*/
	var conf internal.Config
	internal.LoadConfig(&conf)

	source, destination := internal.GetUserInputs(&conf)

	fmt.Println("The source is:", source)
	fmt.Println("The destination is:", destination)
}
