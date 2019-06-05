package gitconfig

// SectionName will be the section header in the .git/config file
const SectionName = "_pair"

// Coauthor represents a coauthor
type Coauthor struct {
	Name  string
	Email string
}

// CurrPairs gets the current co-authors you are pairing with
func CurrPairs() ([]*Coauthor, error) {
	coauthors := []*Coauthor{}

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
