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
func Init(args cli.Args) error {
	fmt.Printf("pair init args: %#v", args)
	gitDir, cmdErr := gitconfig.GitDir()
	gitconfig.CheckCmdError(cmdErr)

	hooksDir := gitDir + "/hooks"
	hookFile := hooksDir + "/prepare-commit-msg"

	if _, err := os.Stat(hookFile); err == nil {
		// Case: path exists & don't reinitialize unless there is the force argument
		if args[0] == "force" {
			_ = makePrepareCommitHook(hookFile)
			fmt.Printf("Forced a reinitialization of the prepare-commit-msg hook\n")
		} else {
			fmt.Printf("Git-pair session is already initialized for this project")
		}
	} else if os.IsNotExist(err) {
		// Case: "prepare-commit-msg" hook does not exit, make one
		fmt.Printf("Prepare-commit-msg does not exist\nMaking one ...")

		_ = makePrepareCommitHook(hookFile)
	}

	return nil
}

func makePrepareCommitHook(filePath string) error {
	// Note: all the file permissions are 755
	fmt.Printf("filepath of makePrepareCommithook: %s", filePath)

	hookScript := []byte(`#!/bin/sh
	set -e

	# Question: Can I alias gitpair command?
	# Hook from git-pair 🍐
	gitpair _prepare-commit-msg $@ #adds all of the arguments in bash
		`)

	err := ioutil.WriteFile(filePath, hookScript, 0755)
	if err != nil {
		log.Fatal(err)
	}

	return nil
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
