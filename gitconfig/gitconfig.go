package gitconfig

import (
	"log"
	"strings"
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

// CurrPairs gets the current co-authors you are pairing with
// func CurrPairs() ([]*Coauthor, error) {
func CurrPairs() ([]string, CmdError) {
	coauthors := []string{}

	exists, err := ContainsSection()
	if err != nil {
		return coauthors, CmdError{}
	}
	if !exists {
		return coauthors, CmdError{}
	}

	output, _ := RunGitConfigCmd("--get-all", "")
	splitOutput := strings.Split(output, "\n")

	return splitOutput, CmdError{}
}

// AddPair adds a new coauthor line only if its unique
// pairStr is in the format of "name <email>"
func AddPair(pairStr string) CmdError {
	// TODO: validate pairStr

	RunGitConfigCmd("--unset-all", pairStr) // Note: prevents the addition of  duplicate keys
	_, cmdErr := RunGitConfigCmd("--add", pairStr)

	return cmdErr
}

// RemovePair removes a single coauthor
func RemovePair(pairStr string) (bool, CmdError) {
	pairsBefore, _ := CurrPairs()

	numWords := len(strings.Split(pairStr, " "))
	if numWords == 1 {
		pairStr = "^" + pairStr + " "
	} else {
		pairStr = "^" + pairStr + "$"
	}

	// TODO: get all pairs, search for pairString
	RunGitConfigCmd("--unset-all", pairStr)
	pairsAfter, _ := CurrPairs()

	// fmt.Printf("%s\n%s", (pairsAfter), (pairsBefore))
	return bool(len(pairsBefore) > len(pairsAfter)), CmdError{}
}

// RemoveAllPairs removes all the coauthors
func RemoveAllPairs() (bool, CmdError) {
	_, cmdErr := RunGitConfigCmd("--unset-all", "")
	return false, cmdErr
}

// CheckError returns a boolean of whether there was an error that should prevent you from proceeding or not
// QUESTION: should we exit the program here?
func CheckCmdError(err CmdError) {
	switch err.ExitCode {
	case 0:
	case 128:
		// Note: what if there are some pair commands we want to run? pair ls?
		log.Fatal(`Cannot run "$pair" command outside of a git repository`)
	default:
		log.Fatalf("Unknown exit code of: %d\n Error message: %s", err.ExitCode, err.Message)
	}
}
