package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

//Flag comment
type Flag interface {
	fmt.Stringer
	Apply(*flag.FlagSet)
	GetName() string
}

func main() {
	app := cli.NewApp() // Create a new app

	app.Flags = []cli.Flag{

		cli.StringFlag{ // add flags with 3 arguments

			Name:  "name",
			Value: "stranger",
			Usage: "your wonderful name",
		},
		cli.IntFlag{
			Name:  "age",
			Value: 0,
			Usage: "your graceful age",
		},
	}

	// function parses and brings data in cli.Context struct
	app.Action = func(c *cli.Context) error {
		log.Printf("Hello %s (%d years), Welcome to the world", c.String("name"), c.Int("age"))
		return nil
	}
	app.Run(os.Args) // Pass os.Args to cli app to parse content
}
