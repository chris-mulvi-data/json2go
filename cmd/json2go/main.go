package main

import (
	"fmt"
	"json2go/internal"
)

func main() {

	source, destination := internal.GetUserInputs()

	fmt.Println("The source is:", source)
	fmt.Println("The destination is:", destination)
}
