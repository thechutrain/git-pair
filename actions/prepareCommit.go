package actions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/thechutrain/git-pair/gitconfig"

	"github.com/thechutrain/git-pair/arrays"
)

// PrepareCommitMsg prepares the commit message
// gets called by git arg [command, "preparecommitmsg", $1, $2, $3]
//  where $1 is the filename of the temp git commit file?
func PrepareCommitMsg(args []string) {
	fileName := args[2]

	lines := readLines(fileName)

	pairs, _ := gitconfig.CurrPairs()
	coauthors := arrays.Map(pairs, func(str string) string {
		return "Co-authored-by: " + str
	})
	updateCommitMsg := addCoAuthors(lines, coauthors)
	writeLines(fileName, updateCommitMsg)

	// TODO: Print out possible co-authors so end-user knows who was added
}

// readLines - given the path to .git/COMMIT_EDITMSG, reads contents into a string slice
func readLines(fileName string) []string {
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

// writeLines - given the path to .git/COMMIT_EDITMSG, will rewrite the contents of the commit msg
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

// addCoAuthors - adds coauthors to the commit message above the first commented block
func addCoAuthors(lines []string, coauthors []string) []string {
	lines = arrays.Filter(lines, func(str string) bool {
		match, _ := regexp.MatchString("^Co-authored-by:", str)
		return !match
	})

	// read from the bottom until there is no comment
	i := 0

	for i = len(lines) - 1; i >= 0; i-- {
		match, _ := regexp.MatchString("^#", lines[i])
		if !match {
			break
		}
	}

	firstLineOfComment := i + 1

	updateCommitMsg := make([]string, 0)
	updateCommitMsg = append(updateCommitMsg, lines[:firstLineOfComment]...)
	updateCommitMsg = append(updateCommitMsg, "# Added by git-pair 🍐")
	updateCommitMsg = append(updateCommitMsg, coauthors...)
	updateCommitMsg = append(updateCommitMsg, "")
	updateCommitMsg = append(updateCommitMsg, lines[firstLineOfComment:]...)

	return updateCommitMsg
}
