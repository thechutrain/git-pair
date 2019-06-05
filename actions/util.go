package actions

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"

	"github.com/pkg/errors"
)

// Coauthor represents a coauthor
type Coauthor struct {
	Name  string
	Email string
}

// containsSection() checks the config file to see if the section exists
func containsSection(filepath string, sectionName string) (bool, error) {
	// TODO: file resolution!
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

// CurrPairs gets the current pairs from .git/config
func CurrPairs() ([]Coauthor, error) {
	coAuthors := []Coauthor{}
	// TODO: get the correct string path where .git/ dir exists
	// exists, _ := containsPairSection("../git-pair/.git/config")
	exists, err := containsSection(".git/config", "pair")
	if err != nil {
		return coAuthors, err
	}

	// Case: if [pair] section does not exist, assume there are no coauthors
	if !exists {
		fmt.Println("[pair] section does not exist. So no coauthors")
		return coAuthors, nil
	}

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("git", []string{"config", "--get-all", "pair.coauthor"}...)
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	// // Note: using teh cmd.Stdout and cmd.Stderr to check for errors
	err = cmd.Run() // Note: subprocess
	if err != nil {
		return coAuthors, errors.Wrap(err, "Failed to execute \"git config --get-all pair.coauthor\" command")
	}

	return coAuthors, nil
}

// createPairSection creates the "pair" section in the .git/config directory so it doesn't throw an error
func addPairSection() error {
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("git", []string{"config", "pair.coauthor"}...)
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return errors.Wrap(err, "Failed to exec \"git config pair.coauthor\"")
	}

	fmt.Println(out.String())
	fmt.Println(stderr.String())

	return nil
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
