package gitconfig

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

// === Helper functions ===

// isNewCoauthor checks in the root config folder to see who the coauthors are
func isNewCoauthor() (bool, error) {
	fmt.Print(SectionName)
	return false, nil
}

// RunCmd a wrapper for exec
func RunCmd(cmdArgs []string) (string, error) {
	if len(cmdArgs) == 0 {
		return "", errors.New("Need at least 1 argument")
	}
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Stderr = &stderr
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		// Question: should we Wrap this error if it will be wrapped where its caught
		return "", errors.Wrap(err, "Failed to run \"RunCmd()\" with arguments "+strings.Join(cmdArgs, ", "))
	}

	return strings.Trim(out.String(), "\n"), nil
}

// GitDir gets the file path to the git dir where the .git/ is located
func GitDir() (string, error) {
	//TODO: change name to getGitDirectory
	// Note:You can use: git rev-parse --git-dir
	out, err := RunCmd([]string{"git", "rev-parse", "--git-dir"})
	isRelativePath := (out == ".git")

	if isRelativePath {
		out, err = RunCmd([]string{"pwd"})
		out = out + "/.git"
	}

	return out, err
}

// ContainsSection checks if pair section exists
func ContainsSection() (bool, error) {
	filepath, err := GitDir()
	if err != nil {
		return false, err
	}
	filepath = filepath + "/config"

	file, err := os.Open(filepath)
	if err != nil {
		errMsg := "Could not open file/filepath: " + filepath
		return false, errors.Wrap(err, errMsg)
	}
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("^\\[" + SectionName + "\\]")
	for scanner.Scan() {
		line := scanner.Text()
		match := re.MatchString(line)
		if match {
			return true, nil
		}
	}
	return false, nil
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
