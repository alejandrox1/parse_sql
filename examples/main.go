package main

import (
	"fmt"

	"github.com/alejandrox1/parse_sql"
)

func main() {

	sqlCommands := parse_sql.SQLCmds("setup.sql")
	for i, cmd := range sqlCommands {
		fmt.Printf("%d) \"%s\"\n", i, cmd)
	}
}
