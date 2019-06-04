package actions

import (
	"bytes"
	"os/exec"
)

// Remove a collaborator that you are currently pairing with
func Remove(index string) {
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("git", []string{"config", "pair"}...)
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	check(err)

	// gitOutput := out.String()
	// for _, line := range gitOutput {
	// 	fmt.Println(line)
	// }
	// fmt.Printf("%#v\n", out.String())

	// TODO: add the filename
	// fmt.Printf("REMOVE WAS CALLED: %s", index)
	// create a slice of all the current pairs
	// check that index is valid
	// create new slice of current pairs (remove with index)
	// overwrite the file with current pairs
}
