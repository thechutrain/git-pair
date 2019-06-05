package actions

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

// TODO: should be a new package?

// CurrPairs gets the current pairs from .git/config
func CurrPairs() ([]string, error) {
	// TODO: make helper func that checks if [pair] exists in .git/config
	currPairs := []string{}

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("git", []string{"config", "--get-all", "pair"}...)
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	// Note: using teh cmd.Stdout and cmd.Stderr to check for errors
	err := cmd.Run() // Note: subprocess

	// Note: git config will throw an error if the "pair" section is missing
	// Case: no current pairs
	// fmt.Println(stderr.String())
	if stderr.String() == "error: key does not contain a section: pair" {
		return nil, currPairs
	}

	// dec := base64.NewDecoder(base64.StdEncoding, &out)
	// io.Copy(os.Stdout, dec)
	return nil, currPairs
}

// createPairSection creates the "pair" section in the .git/config directory so it doesn't throw an error
func createPairSection() {
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("git", []string{"config", "pair.coauthor"}...)
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if len(stderr.String()) != 0 {
		fmt.Println("stderr was true -- before")
	}

	if len(out.String()) != 0 {
		fmt.Println("stderr was true -- before ")
	}

	err := cmd.Run()
	fmt.Println(stderr.String())

	if len(stderr.String()) != 0 {
		fmt.Println("stderr was true ")
	}

	if len(out.String()) != 0 {
		fmt.Println("stderr was true ")
	}
	check(err)

	dec := base64.NewDecoder(base64.StdEncoding, &out)
	io.Copy(os.Stdout, dec)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
