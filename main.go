package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "drybq"
	app.Usage = "simple command for bq dry-run"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello Kotori-chan!")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
