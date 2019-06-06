package actions

import (
	"bufio"
	"fmt"
	"os"
)

// With - pair with command
// func With(args cli.Args) {
// 	// NOTE: cli.Args combines everything after the flag to be a single argument
// 	fullArg := strings.Fields(args.First())

// 	var newPair *Coauthor
// 	// TODO: find collaborator if there is only one user
// 	// return a pointer to a collaborator struct or Panic!
// 	// else case:
// 	newPair = &Coauthor{
// 		Name:  fullArg[0],
// 		Email: fullArg[1],
// 	}

// 	addPair(newPair)
// }

// func Add()

// AddPair adds a new user who is pairing on the code
func addPair(pair *Coauthor) {
	fmt.Printf("Coauthor: %#v\n", pair)

	// TODO: update where this is looking for current pair
	pairExists := isPairing("current_pairs.txt", pair)
	if pairExists {
		return
	}

	// TODO: add the project dir/ first
	f, err := os.OpenFile("current_pairs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)

	_, err = f.Write([]byte(pair.Name + " " + pair.Email + "\n"))
	check(err)

	err = f.Close()
	check(err)
}

// isPairing checks to see if a collaborator is pairing or not
func isPairing(filename string, pair *Coauthor) bool {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0644)
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	check(err)

	pairStr := pair.Name + " " + pair.Email
	for scanner.Scan() {
		userStr := scanner.Text()
		if userStr == pairStr {
			return true
		}
	}

	return false
}
