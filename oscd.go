// oscd (One-Shot change directory) changes to a target directory to
// execute one command then snaps back.
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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

	// second arg onward is command to run
	command := os.Args[2]

	// change to directory
	err := os.Chdir(targetDir)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	// execute command
	cmd := exec.Command(command, os.Args[3:]...)
	if err = cmd.Run(); err != nil {
		log.Fatalf("error: %v\n", err)
	}

	// change directory back to original working directory

	// TIL: The working directory of every process is process-private,
	// so storing/manually changing back is super unnecessary.
}
