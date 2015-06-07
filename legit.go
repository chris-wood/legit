package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"container/list"
	"encoding/json"
)

type GitCommit struct { 
	Id string `json:"id"`
	Merge string `json:"merge"`
	Author string `json:"author"`
	Date string `json:"date"`
	Comment string `json:"comment"`
}

// func (commit GitCommit) ToString() string {
//     return "TODO"
// }

// func buildGitCommit(details []string) (GitCommit) {
// 	commit := GitCommit{id: "0x123"}
// 	return commit 
// }

func linearizeLog(login []byte) {
	fmt.Println("Parsing the log...");

	lines := strings.Split(string(login), "\n")
	commits := list.New()

	commit := ""
	merge := ""
	author := ""
	date := ""
	comment := ""
	parsedComment := false

	for _,line := range(lines) {
		
		// Parse the commit 
		if strings.Contains(line, "commit") {
			commit = line
		} else if strings.Contains(line, "Merge") {
			merge = line
		} else if strings.Contains(line, "Author") {
			author = line
		} else if strings.Contains(line, "Data") {
			date = line
		} else if len(line) > 0 {
			comment = line
		} else if parsedComment == false && len(line) == 0 {
			parsedComment = true
		}

		if parsedComment == true {
			newCommit := GitCommit{Id: commit, Merge: merge, Author: author, Date: date, Comment: comment}
			commits.PushFront(newCommit)

			jsonRep, _ := json.Marshal(newCommit)
			fmt.Println(string(jsonRep))

			commit = ""
			merge = ""
			author = ""
			date = ""
			comment = ""
			parsedComment = false
			break
		}
	}

	// return commits
}

func main() {
	fmt.Println("Hello, there... let's make git better.")	

	lsout, _ := exec.Command("ls").Output()
	fmt.Printf("Contents: %s", lsout)

	args := os.Args[1:]

	// TODO: parse commands
	// this is where we'd use NLP to make this a more user-friendly interface

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

	linearizeLog(logout)

	// Clean up, return to the current directory
	err = os.Chdir(cwd)
}
