package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Hello, there... let's make git better.")	

	lsout, _ := exec.Command("ls").Output()
	fmt.Printf("Contents: %s", lsout)

	args := os.Args[1:]

	// TODO: parse commands

	cwd, err := os.Getwd()
	fmt.Printf("Current working directory: %s\n", cwd)
	if err != nil {
		fmt.Println("Could not get the current working directory")
		os.Exit(-1)
	}

	err = os.Chdir(args[0])
	if err != nil {
		fmt.Println("Error: couldn't change directories")
		os.Exit(-1)
	}

	logout, err := exec.Command("git", "log").Output()
	if err != nil {
		fmt.Printf("Error: failed to run `git log`: %s\n", err)
		os.Exit(-1)
	}

	fmt.Printf("%s\n", logout)

	// Clean up, return to the current directory
	err = os.Chdir(cwd)
}
