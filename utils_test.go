package setup_sqldb

import (
	"reflect"
	"testing"
)

func Test_delEmptyLines(t *testing.T) {
	var tests = []struct {
		s, want []string
	}{
		{[]string{""}, nil},
		{[]string{"rm", "", "one"}, []string{"rm", "one"}},
		{[]string{"", "", "a"}, []string{"a"}},
	}

	for _, test := range tests {
		got := delEmptyLines(test.s)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Delete Empty lines(%q) == %q, want %q", test.s, got, test.want)
		}
	}
}

func Test_trimSQLCmds(t *testing.T) {
	var tests = []struct {
		s    string
		want []string
	}{
		{"", nil},
		{"drop table posts cascade if exists;\n", []string{"drop table posts cascade if exists;"}},
		{"\ncreate table comments ( \tid serial prymary key,\n\tcontent text,\n\tauthor carchar(255),\n\tpost_id integer references posts(id)\n);\n",
			[]string{"create table comments ( id serial prymary key, content text, author carchar(255), post_id integer references posts(id) );"}},
		{`                                                                                
		create table posts (
			id serial primary key,                                                      
			content text,                                                               
			author varchar(255)                                                         
		);`, []string{"create table posts ( id serial primary key, content text, author varchar(255) );"}},
	}

	for _, test := range tests {
		got := trimSQLCmds(test.s)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Delete Empty lines(%q) == %q, want %q", test.s, got, test.want)
		}
	}
}
