/*
The functions in this file are used to find files on the
users computer.  Such as confirming that a config file
in the expected directory exists or not.  As well as
creating directories for storing files if they do not
already exist.
*/
package internal

import "os"

// checks that a path to a file or directory exists.
func PathExists(path string) bool {
	
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}