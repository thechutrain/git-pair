package actions

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/thechutrain/git-pair/gitconfig"
	"gopkg.in/urfave/cli.v1"
)

// Init - creates the prepare-commit-msg hooks
func Init() error {
	gitDir, cmdErr := gitconfig.GitDir()
	gitconfig.CheckCmdError(cmdErr)

	hooksDir := gitDir + "/hooks"
	hookFile := hooksDir + "/prepare-commit-msg"

	if _, err := os.Stat(hookFile); err == nil {
		// path exists; alert user that its been initialized?
		// TODO: if given the --force flag rewrite it, else print that already has a
		// prepare-commit-msg hook
		fmt.Println("File prepare-commit-msg EXISTS")
	} else if os.IsNotExist(err) {
		// file does not exist; make the new script
		fmt.Println("Prepare-commit-msg does not exist")

	}

	makePrepareCommitHook(hookFile)

	return nil
}

func makePrepareCommitHook(filePath string) {
	// Note: all the file permissions are 755
	fmt.Printf("filepath of makePrepareCommithook: %s", filePath)

	hookScript := []byte(`#!/bin/sh
	set -e

	# Question: Can I alias gitpair command?
	# Hook from git-pair 🍐
	gitpair _prepare-commit-msg $@ #adds all of the arguments in bash
		`)

	fmt.Println(hookScript)

	err := ioutil.WriteFile(filePath, hookScript, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

// Add - adds a new pair
func Add(args cli.Args) error {
	return gitconfig.AddPair(strings.Join(args, " "))
}

// Remove - remove
func Remove(args cli.Args) error {
	_, cmdErr := gitconfig.RemovePair(strings.Join(args, " "))
	gitconfig.CheckCmdError(cmdErr)

	return cmdErr
}

// RemoveAll - removes all
func RemoveAll(args cli.Args) error {
	_, cmdErr := gitconfig.RemoveAllPairs()
	gitconfig.CheckCmdError(cmdErr)
	return cmdErr
}

// Status - status status status status status status status status status status status
func Status(args cli.Args) error {
	pairs, _ := gitconfig.CurrPairs()

	if len(pairs) > 0 {
		// TODO: feature - print with column headers etc
		fmt.Printf("Pairing with: \n\t" + strings.Join(pairs, "\n\t"))
		fmt.Printf("\n\nType: \"git remove [name]\" ")
	} else {
		fmt.Printf("You are not currently pairing with anyone\nTo begin pairing with a new person type:\n\t\"git pair add [name]\"\n")
	}

	return nil
}

// ================ HELPER FUNCTIONS ===============
// checkInit - checks that the project has a prepare-commit-msg hook, that its been initialized
func checkInit() error {
	return nil
}
