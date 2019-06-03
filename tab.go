package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/thechutrain/tabcompletion/actions"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	// if args is prepare-commit-msg
	prepareCommit := len(os.Args) > 1 && os.Args[1] == "prepare-commit-msg"
	if prepareCommit {
		fmt.Println("YO0000u are going to prepare commit msg!!!")
		actions.PrepareCommitMsg(os.Args[2:])
		return
	}

	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name: "with",
			// Aliases: []string{"w"},
			Usage: "username email",
			Action: func(c *cli.Context) error {
				actions.With(c.Args())

				return nil
			},
			BashComplete: func(c *cli.Context) {
				if c.NArg() > 0 {
					return
				}

				for _, user := range getCollaborators() {
					fmt.Println(user)
				}
			},
		},
		{
			Name: "remove",
			Action: func(c *cli.Context) error {
				actions.Remove(c.Args().First())
				return nil
			},
		},
		{
			Name: "init",
			Action: func(c *cli.Context) error {
				err := actions.Init()
				return errors.Wrap(err, "could not initialize a new pear project")
			},
		},
		{
			Name: "status",
			Action: func(c *cli.Context) error {
				actions.Status()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getCollaborators() []string {
	file, err := os.Open("collaborators.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
