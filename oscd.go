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
			log.Fatalf("error: %v\n", err)
		}
		targetDir = td
	}

	// second arg onward is command to run
	command := os.Args[2]

	// check if binary exists
	if _, err := exec.LookPath(command); err != nil {
		log.Fatalf("error: %v\n", err)
	}

	// build command
	cmd := exec.Command(command, os.Args[3:]...)
	cmd.Dir = targetDir
	cmd.Env = os.Environ()

	// execute command
	if err := cmd.Run(); err != nil {
		log.Fatalf("error: %v\n", err)
	}

	// TIL: The working directory of every process is process-private,
	// so storing/manually changing back is super unnecessary.
}
