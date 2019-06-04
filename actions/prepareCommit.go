package actions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

// PrepareCommitMsg prepares the commit message
func PrepareCommitMsg(fileName string) {
	lines := readLines(fileName)

	if !containsCoAuthor(lines) {
		coauthors := []string{"Co-authored-by: foobar"}
		updateCommitMsg := addCoAuthors(lines, coauthors)
		writeLines(fileName, updateCommitMsg)
	}

	// TODO: Print out possible co-authors so end-user knows who was added
}

func readLines(fileName string) []string {
	// Open the COMMIT_EDITMSG, and read the content of the whole file into a slice.
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("Failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	defer file.Close()

	return lines
}

func writeLines(fileName string, lines []string) {
	fileWrite, err := os.OpenFile(fileName, os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("Failed opening file: %s", err)
	}

	for _, line := range lines {
		fileWrite.WriteString(line + "\n")
		fmt.Println(line)
	}
	defer fileWrite.Close()
}

func containsCoAuthor(lines []string) bool {
	for _, line := range lines {
		match, _ := regexp.MatchString("(?i)^Co-authored-by:", line)
		if match {
			return true
		}
	}
	return false
}

func addCoAuthors(lines []string, coauthors []string) []string {
	// read from the bottom until there is no comment
	i := 0

	for i = len(lines) - 1; i >= 0; i-- {
		match, _ := regexp.MatchString("^#", lines[i])
		if !match {
			break
		}
	}

	firstLineOfComment := i + 1

	// TODO: make a function that will get current pairs, input will be project name
	updateCommitMsg := make([]string, 0)

	updateCommitMsg = append(updateCommitMsg, lines[:firstLineOfComment]...)
	updateCommitMsg = append(updateCommitMsg, "# Added by 🐙")
	updateCommitMsg = append(updateCommitMsg, coauthors...)
	updateCommitMsg = append(updateCommitMsg, "")
	updateCommitMsg = append(updateCommitMsg, lines[firstLineOfComment:]...)

	return updateCommitMsg
}
