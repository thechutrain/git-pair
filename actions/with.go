package actions

import "fmt"

// AddPair adds a new user who is pairing on the code
func AddPair(name string, email string) {
	fmt.Printf("pairing with: %#v, %#v\n", name, email)
}
