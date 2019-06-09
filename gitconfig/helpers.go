package gitconfig

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

// SectionName will be the section header in the .git/config file
const sectionName = "pair"
const sectionKey = sectionName + "." + "coauthor"

// TODO: make a helper that converts a string of usernames & emails to coauthor structs
// also should validate the strings
// for _, line := range splitOutput {
// 	lineSlice := strings.Split(line, " ")
// 	coauthor := Coauthor{Name: lineSlice[0], Email: lineSlice[1]}
// 	coauthors = append(coauthors, &coauthor)
// }

// ContainsSection checks if pair section exists
func ContainsSection() (bool, error) {
	filepath, cmderr := GitDir()
	// Question: now that Im return error instead of cmderrors
	// can't just pass a regular error, need to pass a CmdError type
	if CheckCmdError(cmderr) != nil {
		return false, cmderr
	}

	filepath = filepath + "/config"

	file, err := os.Open(filepath)
	if err != nil {
		errMsg := "Could not open file/filepath: " + filepath
		return false, errors.Wrap(err, errMsg)
	}
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile("^\\[" + sectionName + "\\]")
	for scanner.Scan() {
		line := scanner.Text()
		match := re.MatchString(line)
		if match {
			return true, nil
		}
	}
	return false, nil
}

// GitDir gets the file path to the where the git dir is located
func GitDir() (string, error) {
	out, cmderr := RootDir()
	if CheckCmdError(cmderr) != nil {
		return "", cmderr
	}
	out = out + "/.git"

	return out, nil
}

func RootDir() (string, error) {
	//TODO: change name to getGitDirectory
	// Note:You can use: git rev-parse --git-dir
	out, err := RunCmd([]string{"git", "rev-parse", "--git-dir"})
	isRelativePath := (out == ".git")

	if isRelativePath {
		// TODO: can use the filepath package to get dir instead
		out, err = RunCmd([]string{"pwd"})
	}

	return out, err
}

// RunGitConfigCmd - Executes "git config" related commands
func RunGitConfigCmd(flags string, val string) (string, error) {
	return RunCmd([]string{"git", "config", flags, sectionKey, val})
}

// RunCmd a wrapper for exec
func RunCmd(cmdArgs []string) (string, error) {
	if len(cmdArgs) == 0 {
		return "", CmdError{Message: "Need at least one argument to run cmd"}
	}

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Stderr = &stderr
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		cmdErr := CmdError{
			Message:  "Failed to run \"RunCmd()\" with arguments " + strings.Join(cmdArgs, ", "),
			ExitCode: exitCode(err)}

		return "", cmdErr
	}

	return strings.Trim(out.String(), "\n"), nil
}

// gets the exit code from a exec.Cmd
func exitCode(err error) int {
	if exitError, ok := err.(*exec.ExitError); ok {
		return exitError.ExitCode()
	}
	return 0
}
