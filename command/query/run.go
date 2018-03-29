package query

import (
	"fmt"
	"regexp"
	"strconv"

	"bufio"
	"os"

	"github.com/mattn/go-shellwords"
	"github.com/snackmgmg/drybq/utils"
	"github.com/urfave/cli"
)

// ToDo: read from config file
const COSTPERTB = 5.0

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
		queryBytes, err := getQueryBytes(string(out))
		if err != nil {
			return err
		}
		cost, err := getCost(queryBytes)
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

func getQueryBytes(str string) (string, error) {
	regex := regexp.MustCompile(`running this query will process (\d+) bytes of data.`)
	queryBytes := regex.FindStringSubmatch(string(str))
	if len(queryBytes) != 2 {
		return "", fmt.Errorf("unexpected result: bytes count is %d, must be %d", len(queryBytes), 1)
	}
	return queryBytes[1], nil
}

func getCost(size string) (float64, error) {
	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		return 0.0, err
	}
	tByte := convertByteToTByte(float64(sizeInt))
	return COSTPERTB * tByte, nil
}

func convertByteToTByte(b float64) float64 {
	return b / (1024 * 1024 * 1024 * 1024)
}
