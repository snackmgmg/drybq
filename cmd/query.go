package cmd

import (
	"fmt"

	"bufio"
	"os"

	"github.com/mattn/go-shellwords"
	"github.com/snackmgmg/drybq/utils"
	"github.com/urfave/cli"
)

var dryBase = "bq query --dry_run %s"
var forceBase = "bq query %s"

func Query(c *cli.Context) error {
	args := utils.CombineStrings(c.Args(), "")
	err := queryExecute(args, true)

	if c.Bool("try") {
		if c.Bool("force") {
			err = queryExecute(args, false)
		} else {
			stdin := bufio.NewScanner(os.Stdin)
			print("execute this query?(Y/N): ")
			stdin.Scan()
			if stdin.Text() == "Y" || stdin.Text() == "y" {
				err = queryExecute(args, false)
			}
		}
	}

	return err
}

func queryExecute(args string, isDry bool) error {
	base := forceBase
	if isDry {
		base = dryBase
	}
	origin := fmt.Sprintf(base, args)
	parsed, err := shellwords.Parse(origin)
	cmd, err := utils.MakeCmd(parsed)
	if err != nil {
		return err
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v: %s", err, string(out))
	}
	result := ""
	if isDry {
		queryBytes, err := utils.GetQueryBytes(string(out))
		if err != nil {
			return err
		}
		cost, err := utils.GetCost(queryBytes)
		if err != nil {
			return err
		}
		result = fmt.Sprintf("%sEstimated Cost is $%s ", string(out), fmt.Sprint(cost))

	} else {
		result = fmt.Sprintf("%s", string(out))
	}
	fmt.Printf("%s\n", result)
	return nil
}
