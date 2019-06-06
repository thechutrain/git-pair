package gitconfig

import (
	"fmt"
	"strings"
)

// SectionName will be the section header in the .git/config file
const SectionName = "pair"

// RunGitConfigCmd - does things
func RunGitConfigCmd(flags string, val string) (string, error) {
	return RunCmd([]string{"git", "config", flags, "pair.coauthor", val})
}

// Coauthor represents a coauthor
type Coauthor struct {
	Name  string
	Email string
}

// CurrPairs gets the current co-authors you are pairing with
// func CurrPairs() ([]*Coauthor, error) {
func CurrPairs() ([]string, error) {
	coauthors := []string{}

	exists, err := ContainsSection()
	if err != nil || !exists {
		return coauthors, err
	}

	output, err := RunGitConfigCmd("--get-all", "")
	fmt.Printf("%s\n", output)
	splitOutput := strings.Split(output, "\n")

	return splitOutput, nil

	// TODO: make a helper function that takes string of coauthors and makes struct? validates
	// for _, line := range splitOutput {
	// 	lineSlice := strings.Split(line, " ")
	// 	coauthor := Coauthor{Name: lineSlice[0], Email: lineSlice[1]}
	// 	coauthors = append(coauthors, &coauthor)
	// }

	// return coauthors, nil
}

// AddPair adds a new
func AddPair(pairStr string) error {
	// TODO: validate pairStr
	// pairStr ="name <email>"

	RunGitConfigCmd("--unset-all", pairStr)
	_, err := RunGitConfigCmd("--add", pairStr)

	return err
}

// RemovePair removes a single coauthor
func RemovePair(pairStr string) bool {
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

	fmt.Printf("%s\n%s", (pairsAfter), (pairsBefore))
	return bool(len(pairsBefore) > len(pairsAfter))
}

// RemoveAllPairs removes all the coauthors
func RemoveAllPairs() (bool, error) {
	_, err := RunGitConfigCmd("--unset-all", "")
	return false, err
}
