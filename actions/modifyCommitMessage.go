package actions

import (
	"bufio"
	"log"
	"os"
	"regexp"

	"github.com/thechutrain/git-pair/gitconfig"

	"github.com/thechutrain/git-pair/arrays"
)

// PrepareCommitMsg prepares the commit message
// gets called by git arg [command, "preparecommitmsg", $1, $2, $3]
//  where $1 is the filename of the temp git commit file?
func ModifyCommitMsg(args []string) {
	fileName := args[2]

	lines := readLines(fileName)

	pairs, _ := gitconfig.CurrPairs()
	coauthors := arrays.Map(pairs, func(str string) string {
		return "Co-authored-by: " + str + " <" + str + "@users.noreply.github.com>"
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

		// fmt.Println(line)
	}
	defer fileWrite.Close()
}

// addCoAuthors - adds coauthors to the commit message above the first commented block
func addCoAuthors(lines []string, coauthors []string) []string {
	// Note: for gc --amend case. Remove possible stale data
	re := regexp.MustCompile("(^Co-authored-by:)|(^# Added by git-pair)")
	lines = arrays.Filter(lines, func(str string) bool {
		return !re.MatchString(str)
	})

	// No co-authors case
	if len(coauthors) == 0 {
		return lines
	}

	// Read from top and find first line that is a comment
	firstEmptyLine := len(lines) - 1
	for i := 0; i < len(lines); i++ {
		match, _ := regexp.MatchString("^#||^\\s", lines[i])
		if match {
			firstEmptyLine = i + 1
			break
		}
	}

	updateCommitMsg := make([]string, 0)
	updateCommitMsg = append(updateCommitMsg, lines[:firstEmptyLine]...)
	updateCommitMsg = append(updateCommitMsg, "")
	updateCommitMsg = append(updateCommitMsg, coauthors...)
	// TODO: adding the #added by git-pair 🍐 breaks it haha
	// updateCommitMsg = append(updateCommitMsg, "# Added by git-pair 🍐")
	updateCommitMsg = append(updateCommitMsg, lines[firstEmptyLine:]...)

	return updateCommitMsg
}
