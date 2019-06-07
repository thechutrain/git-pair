package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/thechutrain/git-pair/actions"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	// if args is prepare-commit-msg
	prepareCommit := len(os.Args) > 1 && os.Args[1] == "prepare-commit-msg"
	if prepareCommit {
		actions.PrepareCommitMsg(os.Args)
		return
	}

	app := cli.NewApp()
	// app.HelpName := "Hi Im help"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "username <email>",
			Action: func(c *cli.Context) error {
				// actions.With(c.Args())
				actions.Add(c.Args())

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
			Name:    "remove",
			Aliases: []string{"rm"},
			Action: func(c *cli.Context) error {
				actions.Remove(c.Args())
				return nil
			},
		},
		{
			Name:    "stop",
			Aliases: []string{"reset"},
			Action: func(c *cli.Context) error {
				actions.RemoveAll(c.Args())
				return nil
			},
		},
		{
			Name:    "status",
			Aliases: []string{"s"},
			Action: func(c *cli.Context) error {
				actions.Status(c.Args())

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
