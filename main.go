package main

import (
	"os"

	"fmt"

	"github.com/snackmgmg/drybq/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	// command name is `drybq`
	app.Name = "drybq"
	app.Usage = "Simple command for bq dry-run with useful info"
	app.UsageText = "drybq command [command options] [arguments...]"
	app.Version = "0.0.1"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			// subcommand #1: query, this command returns processing byte and cost
			Name:   "query",
			Usage:  "Command for bq dry-run and useful info",
			Action: cmd.Query,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "try, t",
					Usage: "Execute query after checked dry-run result.",
				},
				cli.BoolFlag{
					Name:  "force, f",
					Usage: "Force execute query. must be use with 'try' flag",
				},
			},
		},
		{
			// subcommand #2: bulk, this command returns processing byte and cost from csv to csv
			Name:   "bulk",
			Usage:  "bulk import for csv, and for dry-run with useful info",
			Action: cmd.Bulk,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
}
