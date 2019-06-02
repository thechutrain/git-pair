package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "make an explosive entrance"
	app.Action = func(c *cli.Context) error {
		fmt.Printf("Args: %#v", c.Args())
		fmt.Println("Boom yo!")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
