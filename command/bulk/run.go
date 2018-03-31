package bulk

import (
	"bufio"
	"fmt"
	"os"

	"strings"

	"github.com/urfave/cli"
)

// ToDo: read from config file
const COSTPERTB = 5.0

type QueryInfo struct {
	query string
	exec  string
	size  string
	cost  float64
}

func Run(c *cli.Context) error {
	fname := c.Args().Get(0)
	if fname == "" {
		return fmt.Errorf("must be given file name")
	}
	csv, err := readAll(fname)
	if err != nil {
		return err
	}
	queries := []*QueryInfo{}
	for _, c := range csv {
		columns := strings.Split(c, ",")
		query := columns[0]
		e := fmt.Sprintf("bq query --dry_run %s", query)
		queries = append(queries, &QueryInfo{query: query, exec: e})
	}
	return nil
}

func readAll(filename string) ([]string, error) {
	fp, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf(filename + " can't be opened")
	}

	ans := []string{}

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		ans = append(ans, scanner.Text())
	}

	fp.Close()
	return ans, nil
}
