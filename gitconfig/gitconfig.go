package gitconfig

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
	// pathToConfig :=  ootDir()
	// sectionExists, err := containsSection()
	// check(err)

	// check if there pair section

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
