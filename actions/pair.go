package actions

import (
	"fmt"
	"strings"

	"github.com/thechutrain/git-pair/gitconfig"
	"gopkg.in/urfave/cli.v1"
)

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
