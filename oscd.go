// oscd (One-Shot change directory) changes to a target directory to
// execute one command then snaps back.
package main

import (
	"log"
	"os"
)

// cwd is the directory from which this command was invoked
var cwd string

func init() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func main() {
	// first arg is target directory

	// second arg onward is command to run

	// change to directory

	// execute command

	// change directory back to original working directory

	// exit
}
