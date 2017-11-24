package main

import (
	"fmt"

	"github.com/alejandrox1/setup_sqldb"
)

func main() {

	sqlCommands := setup_sqldb.SQLCmds("setup.sql")
	for i, cmd := range sqlCommands {
		fmt.Printf("%d) \"%s\"\n", i, cmd)
	}
}
