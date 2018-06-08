package main

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
	"github.com/BeyondBankingDays/minions-api/command"
)

func main() {
	app := cli.NewApp()

	app.Name = "Hackathon API"
	app.Usage = "Hackathon API"

	app.Commands = []cli.Command{
		command.Server,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("Command exited with error: %+v\n", err)
	}
}
