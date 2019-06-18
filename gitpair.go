package main

import (
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/thechutrain/git-pair/actions"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	// Note: prepare-commit-msg is a func called internally so we don't want to expose it to the cli
	prepareCommit := len(os.Args) > 1 && os.Args[1] == "_prepare-commit-msg"
	if prepareCommit {
		actions.PrepareCommitMsg(os.Args)
		return
	}

	// Registers the cli package for autocompletion of bash commands
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name: "init",
			Action: func(c *cli.Context) error {
				actions.Init(c.Args())

				return nil
			},
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "GitHubHandle",
			Action: func(c *cli.Context) error {
				err := actions.Add(c.Args())
				if err != nil {
					log.Fatal(errors.Cause(err))
				}

				return nil
			},
			BashComplete: func(c *cli.Context) {
				if c.NArg() > 0 {
					return
				}

				// TODO: get a list of past collaborators! to populate the script
				// Get
				// for _, user := range getCollaborators() {
				// 	fmt.Println(user)
				// }
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
