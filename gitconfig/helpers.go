package gitconfig

import "fmt"

// === Helper functions ===

// isNewCoauthor checks in the root config folder to see who the coauthors are
func isNewCoauthor() (bool, error) {
	fmt.Print(SectionName)
	return false, nil
}

// rootDir gets the root dir where the .git/ is located
func rootDir() string {
	// Note:You can use: git rev-parse --git-dir
	return ""
}

// containsSection
func containsSection(filepath string, sectionName string) (bool, error) {
	return false, nil
}
