package actions

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/urfave/cli.v1"
)

// With - pair with command
func With(args cli.Args) {
	words := strings.Fields(args[0])
	for _, word := range words {
		fmt.Println(word)
	}
	fmt.Printf("length of words: %d\n", len(words))

	// if only one argument, check that we have that user & email
	addPair(words[0], words[1])

	// if two arguments OKAY --> write to current_pair

}

// AddPair adds a new user who is pairing on the code
func addPair(name string, email string) {
	fmt.Printf("pairing with: %#v, %#v\n", name, email)

	//TODO: check that the given pair is not already in the file

	f, err := os.Create("current_pairs.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.WriteString("hi there")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(l)
	defer f.Close()

	// newPair := []byte(name + email)
	// err := ioutil.WriteFile("/currently_pairing.txt", newPair, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
