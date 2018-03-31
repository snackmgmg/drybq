package query

import (
	"fmt"

	"bufio"
	"os"

	"github.com/mattn/go-shellwords"
	"github.com/snackmgmg/drybq/utils"
	"github.com/urfave/cli"
)

func Run(c *cli.Context) error {
	origin := fmt.Sprintf("bq query --dry_run %s", utils.CombineStrings(c.Args(), ""))
	err := execute(origin, true)

	if c.Bool("try") {
		force := fmt.Sprintf("bq query %s", utils.CombineStrings(c.Args(), ""))
		if c.Bool("force") {
			err = execute(force, false)
		} else {
			stdin := bufio.NewScanner(os.Stdin)
			print("execute this query?(Y/N): ")
			stdin.Scan()
			if stdin.Text() == "Y" || stdin.Text() == "y" {
				err = execute(force, false)
			}
		}
	}

	return err
}

func execute(str string, isDry bool) error {
	args, err := shellwords.Parse(str)
	cmd, err := utils.MakeCmd(args)
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
