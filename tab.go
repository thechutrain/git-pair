package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/thechutrain/tabcompletion/actions"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:    "with",
			Aliases: []string{"w"},
			Usage:   "username email",
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
