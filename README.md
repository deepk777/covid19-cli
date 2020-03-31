# covid19-cli
Started as a weekend project, It is a Command Line Tool to get details of stats on Covid cases. 

[![Go Report Card](https://goreportcard.com/badge/github.com/deepk777/covid19-cli?style=flat-square)](https://goreportcard.com/badge/github.com/deepk777/covid19-cli)

**Table of Contents**

- [Installing](#installing)
- [Usage](#usage)
- [Built With](#built-with)

## Installing 

`go get -u github.com/deepk777/covid19-cli/covid`

## Usage
```
$ covid -h

Covid is easy to use CLI tool for live covid stats.

        Run:
        > covid - Get Global count.

        Subcommands
        > covid full - Get Top counts by active cases
        > covid <country> - Get stats for input country

        Example :
        covid India
        +--------+-------+--------+-----------+--------+
        | REGION | CASES | ACTIVE | RECOVERED | DEATHS |
        +--------+-------+--------+-----------+--------+
        | India  |  1251 |   1117 |       102 |     32 |
        +--------+-------+--------+-----------+--------+

Usage:
  covid [flags]

Flags:
      --format string   Format for displaying output. Options are "table" and "json" (default "table")
  -h, --help            help for covid
      --limit int       Limit results to display for command <covid full>. Only applicable for table format. Minimum value is 1. (default 50)

```



## Built With

* [Golang](https://golang.org) - The Google Go Language
* [GoModules](https://github.com/golang/go/wiki/Modules) - Built-In Golang Modules Dependency Manager


