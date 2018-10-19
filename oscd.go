// oscd (One-Shot change directory) changes to a target directory to
// execute one command then snaps back.
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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

	// get abs path if given relative
	if !filepath.IsAbs(targetDir) {
		td, err := filepath.Abs(targetDir)
		if err != nil {
			panic(err)
		}
		targetDir = td
	}

	// second arg onward is command to run
	command := os.Args[2]

	// check if binary exists
	if _, err := exec.LookPath(command); err != nil {
		panic(err)
	}

	// change to directory
	if err = os.Chdir(targetDir); err != nil {
		log.Fatalf("error: %v\n", err)
	}

	// build command
	cmd := exec.Command(command, os.Args[3:]...)
	cmd.Env = os.Environ()

	// execute command
	cmd := exec.Command(command, os.Args[3:]...)
	if err = cmd.Run(); err != nil {
		log.Fatalf("error: %v\n", err)
	}

	// change directory back to original working directory
	 
	// TIL: The working directory of every process is process-private,
	// so storing/manually changing back is super unnecessary.
}
