package setup_sqldb

import (
	"os"
	"testing"
)

func Test_DBDataSource(t *testing.T) {
	var tests = []struct {
		s    map[string]string
		want string
	}{
		{map[string]string{"USER": "user", "PASSWORD": "pass", "HOST": "db", "DBNAME": "db", "SSLMODE": "verify-full"},
			"user=user password=pass host=db dbname=db sslmode=verify-full"},
		{map[string]string{"USER": "user", "PASSWORD": "pass", "SSLMODE": "disable"},
			"user=user password=pass sslmode=disable"},
	}

	for _, test := range tests {
		setTestEnv(test.s)
		got := DBDataSource()
		if got != test.want {
			t.Errorf("Delete Empty lines(%q) == %q, want %q", test.s, got, test.want)
		}
		unsetTestEnv(test.s)
	}
}

func setTestEnv(env map[string]string) {
	for k, v := range env {
		os.Setenv(k, v)
	}
}

func unsetTestEnv(env map[string]string) {
	for k, _ := range env {
		os.Unsetenv(k)
	}
}
