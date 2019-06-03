package actions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/urfave/cli.v1"
)

// Collaborator struct represents a person who is/or has paired
type Collaborator struct {
	GhName string
	Email  string
}

// With - pair with command
func With(args cli.Args) {
	// NOTE: cli.Args combines everything after the flag to be a single argument
	fullArg := strings.Fields(args.First())
	// for _, word := range fullArg {
	// 	fmt.Println(word)
	// }

	var pair *Collaborator
	// TODO: find collaborator if there is only one user
	// return a pointer to a collaborator struct or Panic!
	// else case:
	pair = &Collaborator{
		GhName: fullArg[0],
		Email:  fullArg[1],
	}

	addPair(pair)
}

// AddPair adds a new user who is pairing on the code
func addPair(pair *Collaborator) {
	isPairing("current_pairs.txt", pair)

	fmt.Printf("Collaborator: %#v\n", pair)

	// TODO: add the project dir/ first
	f, err := os.OpenFile("current_pairs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(pair.GhName + " " + pair.Email + "\n")); err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// isPairing checks to see if a collaborator is pairing or not
func isPairing(filename string, pair *Collaborator) bool {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		pairStr := scanner.Text()
		fmt.Println(pairStr)
	}

	return false
}
