package gitconfig

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/thechutrain/git-pair/arrays"
)

// Coauthor represents a coauthor
type Coauthor struct {
	Name  string
	Email string
}

// CmdError - represents an error from making a bash command
type CmdError struct {
	Message  string // error description
	ExitCode int    // exit code
}

func (e CmdError) Error() string {
	return fmt.Sprintf("CmdError: %s", e.Message)
}

// CurrPairs gets the current co-authors you are pairing with
func CurrPairs() ([]string, error) {
	// TODO: return struct of curr pairs?
	coauthors := []string{}

	exists, err := ContainsSection()
	if err != nil {
		return coauthors, err
	}

	if !exists {
		return coauthors, nil
	}

	output, err := RunGitConfigCmd("--get-all", "")
	if err != nil {
		return coauthors, err
	}

	splitOutput := strings.Split(output, "\n")
	splitOutput = arrays.Filter(splitOutput, func(str string) bool {
		return str != ""
	})

	return splitOutput, nil
}

func validatePairStr(rawArgs []string) error {
	// TODO: validate the pairStr
	// IDEA: could validate github handle via http request?
	/* Valid inputs:
	- githubhandle
	- name <email>
	*/

	switch len(rawArgs) {
	case 0:
		return errors.New("Must provide at least one argument for adding a pair")
	case 1:
		if strings.TrimSpace(rawArgs[0]) == "" {
			return errors.New("Must provide at least one argument for adding a pair")
		}
		rawUserName := rawArgs[0]
		// TODO: look up this github username
	// case 2: // Note: for username and email entered manually

	default:

	}

	return nil
}

// AddPair adds a new coauthor line only if its unique
// pairStr is in the format of "name <email>"
func AddPair(rawArgs []string) error {
	// Validate user input
	err := validatePairStr(rawArgs)
	if err != nil {
		return err
	}

	pairStr := strings.Join(rawArgs, " ")
	_, cmdErr := RunGitConfigCmd("--unset-all", pairStr) // Note: prevents the addition of  duplicate keys
	if cmdErr != nil {
		return cmdErr
	}

	_, cmdErr = RunGitConfigCmd("--add", pairStr)
	if CheckCmdError(cmdErr) != nil {
		return cmdErr
	}

	return nil
}

// RemovePair removes a single coauthor
func RemovePair(pairStr string) (bool, error) {
	pairsBefore, cmdErr := CurrPairs()
	if CheckCmdError(cmdErr) != nil {
		return false, cmdErr
	}

	numWords := len(strings.Split(pairStr, " "))
	if numWords == 1 {
		pairStr = "^" + pairStr + " "
	} else {
		pairStr = "^" + pairStr + "$"
	}

	_, cmdErr = RunGitConfigCmd("--unset-all", pairStr)
	CheckCmdError(cmdErr)
	pairsAfter, _ := CurrPairs()

	// fmt.Printf("%s\n%s", (pairsAfter), (pairsBefore))
	return bool(len(pairsBefore) > len(pairsAfter)), CmdError{}
}

// RemoveAllPairs removes all the coauthors
func RemoveAllPairs() error {
	_, cmdErr := RunGitConfigCmd("--unset-all", "")
	return cmdErr
}

// CheckError returns a boolean of whether there was an error that should prevent you from proceeding or not
// QUESTION: should we exit the program here?
// Using this function as middleware
func CheckCmdError(err error) error {
	if err == nil {
		return nil
	}

	if CmdError, ok := err.(CmdError); ok {
		switch CmdError.ExitCode {
		case 0:
			return nil
		case 1:
			// Note - Git errors, there is a section, but no keys.
			return nil
		case 5:
			// Note - Ignore this error: Git config exit code if you try to access a section.key that does not exist
			return nil
		case 128:
			// fmt.Printf("Exited with code: %d\n", CmdError.ExitCode)
			log.Fatal(`Cannot run pair command outside of a git repository`)
		default:
			log.Fatalf("Unknown exit code of: %d\n Error message: %s", CmdError.ExitCode, CmdError.Message)
		}
	}

	return err
}
