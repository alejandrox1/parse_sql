package parse_sql

import (
	"fmt"
	"os"
	"strings"
)

func DBDataSource() string {
	var source []string

	if user := os.Getenv("USER"); user != "" {
		source = append(source, fmt.Sprintf("user=%s", user))
	}
	if pass := os.Getenv("PASSWORD"); pass != "" {
		source = append(source, fmt.Sprintf("password=%s", pass))
	}
	if host := os.Getenv("HOST"); host != "" {
		source = append(source, fmt.Sprintf("host=%s", host))
	}
	if dbname := os.Getenv("DBNAME"); dbname != "" {
		source = append(source, fmt.Sprintf("dbname=%s", dbname))
	}
	if sslmode := os.Getenv("SSLMODE"); sslmode != "" {
		source = append(source, fmt.Sprintf("sslmode=%s", sslmode))
	}

	return strings.Join(source, " ")
}
