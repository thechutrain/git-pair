package main

import (
	"fmt"
	"log"
	"os"

	"github.com/thechutrain/tabcompletion/actions"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	tasks := []string{"clean", "gym", "tan", "laundry"}

	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:    "with",
			Aliases: []string{"w"},
			Usage:   "username email",
			Action: func(c *cli.Context) error {
				fmt.Printf("Action args: %#v\n", c.Args())
				fmt.Println("completed task: ", c.Args().First())

				// TODO: add username and email to users to pair with
				actions.AddPair("merklebros", "patrick")

				return nil
			},
			BashComplete: func(c *cli.Context) {
				// Get a list of users to pair with from a file

				if c.NArg() > 0 {
					return
				}
				for _, t := range tasks {
					fmt.Println(t)
				}
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}


func getCollaborators() []string{} {

} 