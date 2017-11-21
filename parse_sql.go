package parse_sql

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func cleanString(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func delEmptyLines(s []string) []string {
	var newSlice []string
	for _, str := range s {
		if str != "" {
			newSlice = append(newSlice, str)
		}
	}
	return newSlice
}

func trimSQLCmds(s string) []string {
	cmds := strings.SplitAfter(s, ";")
	for i, c := range cmds {
		cmds[i] = cleanString(c)
	}
	return delEmptyLines(cmds)
}

func sqlCmds(filename string) []string {
	sqlSetup, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return trimSQLCmds(string(sqlSetup))
}

func main() {

	sqlCommands := sqlCmds("setup.sql")
	for i, cmd := range sqlCommands {
		fmt.Printf("%d) \"%s\"\n", i, cmd)
	}
}
