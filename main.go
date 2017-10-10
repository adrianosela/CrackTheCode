package main

import (
	"fmt"
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

var buildVersion string //injected at build time

func main() {

	app := cli.NewApp()
	app.Version = buildVersion
	app.EnableBashCompletion = true
	app.Usage = "cli to manage bruteforce server"
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Println("[ERROR] The command provided is not supported: ", command)
		c.App.Run([]string{"help"})
	}

	app.Commands = []cli.Command{
		{
			Name:  "code",
			Usage: "manages the current code",
			Subcommands: []cli.Command{
				{
					Name:   "generate",
					Usage:  "generates and sets a new code on the server",
					Action: GenerateNewCode,
					Flags: []cli.Flag{
						cli.IntFlag{
							Name:  "max",
							Usage: "The maximum number to generate",
						},
					},
				},
				{
					Name:   "crack",
					Usage:  "generates and sets a new code on the server",
					Action: CrackTheCode,
					Flags: []cli.Flag{
						cli.IntFlag{
							Name:  "num",
							Usage: "The number to try",
						},
						cli.BoolFlag{
							Name:  "all",
							Usage: "Try all numbers in range",
						},
					},
				},
				{
					Name:   "cheat",
					Usage:  "gets the value of the current code",
					Action: Cheat,
				},
				{
					Name:   "tries",
					Usage:  "Gets current number of tries to crack the code",
					Action: Tries,
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("\nError: %s \n", err)
		os.Exit(1)
	}
}
