package parse_sql

import (
	"io/ioutil"
	"strings"
)

// Remove any extrawhitespaces, tabs, and newlines from the command.
func cleanString(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// Return a slice of strings with nonempty strings.
func delEmptyLines(s []string) []string {
	var newSlice []string
	for _, str := range s {
		if str != "" {
			newSlice = append(newSlice, str)
		}
	}
	return newSlice
}

// Take sql commands as a single string and return a slice of strings where
// each string is an sql command.
func trimSQLCmds(s string) []string {
	cmds := strings.SplitAfter(s, ";")
	for i, c := range cmds {
		cmds[i] = cleanString(c)
	}
	return delEmptyLines(cmds)
}

// Given a filename containing all sql commands to be executed, return a slice
// of strings with all the commands parsed.
func SQLCmds(filename string) []string {
	sqlSetup, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return trimSQLCmds(string(sqlSetup))
}
