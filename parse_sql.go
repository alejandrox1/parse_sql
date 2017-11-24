package parse_sql

import (
	"database/sql"
	"fmt"
	"io/ioutil"
)

type SQLSetup struct {
	commands []string
}

// Get all sql commands from a filename and store them in SQLSetup struct.
func (s *SQLSetup) ParseCommands(filename string) {
	sqlSetup, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	s.commands = trimSQLCmds(string(sqlSetup))
}

// Get all sql commands and execute them.
func (s *SQLSetup) Init(db *sql.DB, filename string) {
	s.ParseCommands(filename)

	for _, cmd := range s.commands {
		_, err := db.Exec(cmd)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}
}
