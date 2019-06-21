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
		fmt.Printf("Initializing git commit message hook ...\n")
		_ = makeCommitMessageHook()
	} else if len(args) > 0 && args[0] == "force" {
		fmt.Printf("Reinitializing git commit message hook ... \n")
		_ = makeCommitMessageHook()
	} else {
		fmt.Printf("Already initialized. To remake the commit msg hook, type: pair init force")
	}

	return nil
}

// isInitialized checks to see if the prepare-commit-msg hook exists
func isInitialized() bool {
	gitDir, err := gitconfig.GitDir()
	if err != nil {
		log.Fatal(err)
	}

	hookFile := gitDir + "/hooks/commit-msg"

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

func makeCommitMessageHook() error {
	// Note: all the file permissions are 755
	gitDir, err := gitconfig.GitDir()
	if err != nil {
		return err
	}

	hookFile := gitDir + "/hooks/commit-msg"

	hookScript := []byte(`#!/bin/sh
	set -e

	# Hook from git-pair 🍐
	gitpair _modify-commit-msg $@ #adds all of the arguments in bash
		`)

	err = ioutil.WriteFile(hookFile, hookScript, 0755)
	if err != nil {
		return err
	}

	return nil
}

// Add - adds a new pair
func Add(args cli.Args) error {
	mustBeInitialized()

	return gitconfig.AddPair(args)
}

// Remove - remove
func Remove(args cli.Args) error {
	mustBeInitialized()

	_, err := gitconfig.RemovePair(strings.Join(args, " "))

	return err
}

// RemoveAll - removes all
func RemoveAll(args cli.Args) error {
	mustBeInitialized()

	err := gitconfig.RemoveAllPairs()

	return err
}

// Status - status status status status status status status status status status status
func Status(args cli.Args) error {
	mustBeInitialized()

	pairs, _ := gitconfig.CurrPairs()

	if len(pairs) > 0 {
		// TODO: feature - print with column headers etc
		fmt.Printf("Pairing with: \n\t" + strings.Join(pairs, "\n\t"))
		fmt.Printf("\n\nType: \"pair stop\" ")
	} else {
		fmt.Printf("You are not currently pairing with anyone\nTo begin pairing with a new person type:\n\t\"pair add [GitHub_Handle]\"\n")
	}

	return nil
}
