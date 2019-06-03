package actions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Status prints out who you are currently pairing with
func Status() {
	// TODO: get the filename or project name
	// NOTE: possible that the file does not exist and will throw an error
	// data, err := ioutil.ReadFile("current_pairs.txt")
	file, err := os.Open("current_pairs.txt")
	if err != nil {
		fmt.Println("Your are not currently pairing with anyone")
		return
	}

	var pairs []*Collaborator
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		pairs = append(pairs, &Collaborator{GhName: words[0], Email: words[1]})
	}

	switch len(pairs) {
	case 0:
		fmt.Println("You are not currently pairing with anyone")
	case 1:
		fmt.Println("You are currently pairing with 1 person:")
	default:
		fmt.Printf("You are currenlty pairing with %d people:\n", len(pairs))
	}

	for i, currPair := range pairs {
		fmt.Printf("(%d) %s\n", i, currPair.GhName)
	}

	if len(pairs) > 0 {
		fmt.Printf("To remove a collaborator enter: \"pair remove [index]\"\n")
	}
}
