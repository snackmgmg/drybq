# drybq

[![CircleCI](https://circleci.com/gh/snackmgmg/drybq/tree/master.svg?style=svg)](https://circleci.com/gh/snackmgmg/drybq/tree/master)

Simple and Useful CLI-Tools for BigQuery.

# Installation

This command require bq command-line tool.


```
go get github.com/snackmgmg/drybq
```

# Usage

## query

This command is a wrapper for dry_run.
Get the dry_run result and estimate cost of BigQuery.

```
drybq query "[some query]"
```

### Flags

#### --try, -t

If use this flag, input "Y" and can execute same query and get result.

```
drybq query --try "[some query]"
execute this query?(Y/N): y
[result]
```

#### --force, -f

If use this flag with `--try`, can execute same query and get result without input.

```
drybq query --try "[some query]"
[result]
```

## bulk

This command is dryrun for many queries.
Get cost and size using the csv file with queries.

### csv format

```
[query1]
[query2]
[query3]
```

### command

```
drybq bulk queries.csv > result.csv
```
