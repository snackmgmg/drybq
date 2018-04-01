package cmd

import (
	"bufio"
	"fmt"
	"os"

	"strings"

	"github.com/mattn/go-shellwords"
	"github.com/snackmgmg/drybq/utils"
	"github.com/urfave/cli"
)

type QueryInfo struct {
	query  string
	exec   string
	size   string
	cost   float64
	result string
}

func Bulk(c *cli.Context) error {
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
	err = bulkExecute(queries)
	if err != nil {
		return err
	}
	header := "query, size, cost, result"
	body := makeCSV(queries)
	fmt.Fprintln(os.Stdout, header)
	fmt.Fprint(os.Stdout, body)
	return nil
}

func makeCSV(queries []*QueryInfo) string {
	result := ""
	//ToDo: make reading format from argument
	format := "%s, %s, %f, %s\n"
	for _, q := range queries {
		result += fmt.Sprintf(format, q.exec, q.size, q.cost, q.result)
	}
	return result
}

func bulkExecute(queries []*QueryInfo) error {
	if len(queries) == 0 {
		return fmt.Errorf("don't exist target queries")
	}
	for i, q := range queries {
		args, err := shellwords.Parse(q.exec)
		if err != nil {
			return fmt.Errorf("line %d, %v", i+1, err)
		}
		cmd, err := utils.MakeCmd(args)
		if err != nil {
			return fmt.Errorf("line %d, %v", i+1, err)
		}
		out, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("line %d, %v", i+1, err)
		}
		result := strings.TrimRight(string(out), "\n")
		result = strings.Replace(result, ",", " ", -1)
		queryBytes, err := utils.GetQueryBytes(result)
		if err != nil {
			return fmt.Errorf("line %d, %v", i+1, err)
		}
		cost, err := utils.GetCost(queryBytes)
		if err != nil {
			return fmt.Errorf("line %d, %v", i+1, err)
		}
		q.size = queryBytes
		q.cost = cost
		q.result = result
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
