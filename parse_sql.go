package parse_sql

import (
	"database/sql"
	"io/ioutil"
	"time"
)

type SQLSetup struct {
	DriverName string
	SQLScript  string
	Commands   []string
}

// Get all sql commands from a filename and store them in SQLSetup struct.
func (s *SQLSetup) ParseCommands() {
	sqlSetup, err := ioutil.ReadFile(s.SQLScript)
	if err != nil {
		panic(err)
	}
	s.Commands = trimSQLCmds(string(sqlSetup))
}

// Get all sql commands and execute them.
func (s *SQLSetup) Init(maxTries ...int) (Db *sql.DB, err error) {
	// Create worker pool
	Db, err = sql.Open(s.DriverName, DBDataSource())
	if err != nil {
		return
	}

	// Exponential retry
	tries := 5
	delay := time.Duration(500)
	if len(maxTries) > 0 {
		tries = maxTries[0]
	}
	for ; tries >= 0; tries, delay = tries-1, delay*2 {
		if err = Db.Ping(); err == nil {
			break
		} else if err != nil && tries == 0 {
			return
		}
		time.Sleep(delay * time.Millisecond)
	}

	// Execute sql commands
	s.ParseCommands()
	for _, cmd := range s.Commands {
		_, err = Db.Exec(cmd)
		if err != nil {
			return
		}
	}
	return
}
