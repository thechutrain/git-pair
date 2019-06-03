package actions

import (
	"fmt"
	"io/ioutil"
)

// Status prints out who you are currently pairing with
func Status() {
	// TODO: get the filename or project name
	// NOTE: possible that the file does not exist and will throw an error
	data, err := ioutil.ReadFile("current_pairs.txt")

	// Note: assuming if there was an error opening the file, file doesnt exist
	// perhaps there is a better way to do this?h
	if err != nil {
		fmt.Println("You are not currently pairing with anyone")
	}

	pairs := string(data)
	fmt.Printf("Pairs from with: %#v", pairs)
	// open current_pair
	// log each of those pairs
}
