package main

import (
	"database/sql"

	"github.com/alejandrox1/parse_sql"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	sqlSetup := parse_sql.SQLSetup{
		DriverName: "postgres",
		SQLScript: "setup.sql",
	}
	Db, err = sqlSetup.Init()
	if err != nil {
		panic(err)
	}
}

func (p *Post) Create() (err error) {
	err = Db.QueryRow("insert into posts (content, author) values ($1, $2) returning id",
		p.Content, p.Author).Scan(&p.Id)
	return
}

func (p *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author= $3 where id = $1",
		p.Id, p.Content, p.Author)
	return
}

func (p *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", p.Id)
	return
}

func retrieve(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1",
		id).Scan(&post.Id, &post.Content, &post.Author)
	return
}
