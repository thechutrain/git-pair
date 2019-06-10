package main

import (
	"log"
	"os"

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
			Usage:   "username <email>",
			Action: func(c *cli.Context) error {
				actions.Add(c.Args())

				return nil
			},
			BashComplete: func(c *cli.Context) {
				if c.NArg() > 0 {
					return
				}

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
