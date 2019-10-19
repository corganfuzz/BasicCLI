package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "save",
			Value: "no",
			Usage: "Should save to DB (yes/no)",
		},
	}

	app.Version = "1.0"

	// Define Action

	app.Action = func(c *cli.Context) error {
		var args []string
		if c.NArg() > 0 {
			args = c.Args()
			personName := args[0]
			marks := args[1:len(args)]
			log.Println("Person: ", personName)
			log.Println("marks", marks)
		}

		// Check flag value

		if c.String("save") == "no" {
			log.Println("Skipping saving to the DB")
		} else {
			// DB logic here
			log.Println("Saving to DB", args)
		}
		return nil
	}

	app.Run(os.Args)
}
