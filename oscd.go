// oscd (One-Shot change directory) changes to a target directory to
// execute one command then snaps back.
package main

import (
	"fmt"
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
	if len(os.Args) < 3 {
		// TODO: properly format usage.
		fmt.Println("usage...\n\toscd <target directory> <command>")
		os.Exit(2)
	}

	targetDir := os.Args[1]
	fmt.Printf("Target Directory: %s\n", targetDir)

	// second arg onward is command to run
	command := os.Args[2]
	fmt.Printf("Command: %s\n", command)

	// change to directory

	// execute command

	// change directory back to original working directory

	// exit
}
