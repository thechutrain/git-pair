package actions

import (
	"fmt"
	"strings"

	"github.com/thechutrain/git-pair/gitconfig"
	"gopkg.in/urfave/cli.v1"
)

// Add - adds a new pair
func Add(args cli.Args) error {
	cmdErr := gitconfig.AddPair(strings.Join(args, " "))
	gitconfig.CheckCmdError(cmdErr)
	return nil
}

// Remove - remove
func Remove(args cli.Args) bool {
	//TODO: all these func should return CmdErr
	_, cmdErr := gitconfig.RemovePair(strings.Join(args, " "))
	gitconfig.CheckCmdError(cmdErr)

	return true
}

// RemoveAll - removes all
func RemoveAll(args cli.Args) error {
	_, cmdErr := gitconfig.RemoveAllPairs()
	gitconfig.CheckCmdError(cmdErr)
	return nil
}

// Status - status status status status status status status status status status status
func Status(args cli.Args) {
	pairs, cmdErr := gitconfig.CurrPairs()
	gitconfig.CheckCmdError(cmdErr)

	if pairs != nil {
		// TODO: feature - print with column headers etc
		fmt.Printf("Pairing with: \n\t" + strings.Join(pairs, "\n\t"))
		fmt.Printf("\n\nType: \"git remove [name]\" ")
	}
}
