package gitconfig

import (
	"strings"
)

// SectionName will be the section header in the .git/config file
const SectionName = "pair"

// Coauthor represents a coauthor
type Coauthor struct {
	Name  string
	Email string
}

// CurrPairs gets the current co-authors you are pairing with
func CurrPairs() ([]*Coauthor, error) {
	coauthors := []*Coauthor{}

	exists, err := ContainsSection()
	if err != nil {
		return coauthors, err
	}

	if !exists {
		return coauthors, nil
	}

	// get git config --get-all pair.coauthor
	output, err := RunCmd([]string{"git", "config", "--get-all", SectionName + ".coauthor"})
	splitOutput := strings.Split(output, "\n")
	// ["anush a@email.com"]
	// [&Coauthor{Name: "anush"}]

	for _, line := range splitOutput {
		lineSlice := strings.Split(line, " ")
		//TODO: remove brackets (if any) from email
		// TODO: check line valid input
		coauthor := Coauthor{Name: lineSlice[0], Email: lineSlice[1]}
		coauthors = append(coauthors, &coauthor)

	}
	return coauthors, nil
}

// AddPair adds a new coauthor if they are not currently listed
func AddPair(coauthor *Coauthor) (bool, error) {
	return false, nil
}

// RemovePair removes a single coauthor
func RemovePair(coauthor *Coauthor) (bool, error) {
	return false, nil
}

// RemoveAllPairs removes all the coauthors
func RemoveAllPairs() (bool, error) {
	return false, nil
}
