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
	if !isInitialized() {
		fmt.Printf("Initializing git prepare commit message hook ...\n")
		_ = makePrepareCommitHook()
	} else if len(args) > 0 && args[0] == "force" {
		fmt.Printf("Reinitializing git prepare commit message hook ... \n")
		_ = makePrepareCommitHook()
	} else {
		fmt.Printf("Already initialized. To remake the prepare commit msg hook, type: pair init force")
	}

	return nil
}

// isInitialized checks to see if the prepare-commit-msg hook exists
func isInitialized() bool {
	gitDir, cmdErr := gitconfig.GitDir()
	gitconfig.CheckCmdError(cmdErr)

	hookFile := gitDir + "/hooks/prepare-commit-msg"

	if _, err := os.Stat(hookFile); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else if err != nil {
		log.Fatal(err)
	}

	return false
}

// mustBeInitialized - makes sures the prepare commit msg hook exists
func mustBeInitialized() {
	if !isInitialized() {
		// Question: Panic here? or return error?
		// log.Fatalf("You must first initialize this repo for pairing\nType: git pair init\n")
		fmt.Printf("You must first initialize this repo for pairing\nType: pair init\n")
		os.Exit(2) // TODO: have a list of exit codes?
	}
}

func makePrepareCommitHook() error {
	// Note: all the file permissions are 755
	gitDir, cmdErr := gitconfig.GitDir()
	gitconfig.CheckCmdError(cmdErr)

	hookFile := gitDir + "/hooks/prepare-commit-msg"

	hookScript := []byte(`#!/bin/sh
	set -e

	# Question: Can I alias gitpair command?
	# Hook from git-pair 🍐
	gitpair _prepare-commit-msg $@ #adds all of the arguments in bash
		`)

	err := ioutil.WriteFile(hookFile, hookScript, 0755)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// Add - adds a new pair
func Add(args cli.Args) error {
	mustBeInitialized()

	return gitconfig.AddPair(strings.Join(args, " "))
}

// Remove - remove
func Remove(args cli.Args) error {
	mustBeInitialized()

	_, cmdErr := gitconfig.RemovePair(strings.Join(args, " "))
	gitconfig.CheckCmdError(cmdErr)

	return cmdErr
}

// RemoveAll - removes all
func RemoveAll(args cli.Args) error {
	mustBeInitialized()

	_, cmdErr := gitconfig.RemoveAllPairs()
	gitconfig.CheckCmdError(cmdErr)
	return cmdErr
}

// Status - status status status status status status status status status status status
func Status(args cli.Args) error {
	mustBeInitialized()

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
