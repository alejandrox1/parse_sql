package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/alejandrox1/parse_sql"
	_ "github.com/lib/pq"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func (p *Post) Create() (err error) {
	stmt, err := Db.Prepare("insert into posts (content, author) values ($1, $2) returning id")
	if err != nil {
		fmt.Println("wadahell")
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(p.Content, p.Author).Scan(&p.Id)
	return
}

func (p *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", p.Id, p.Content, p.Author)
	return
}

func (p *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", p.Id)
	return
}

func GetPosts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		return
	}

	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "host=db user=user password=pass dbname=db sslmode=disable")
	if err != nil {
		panic(err)
	}

	maxTries := 5
	delay := time.Duration(100)
	for ; maxTries>=0; maxTries, delay = maxTries-1, delay*2  {

		if err = Db.Ping(); err == nil {
			break
		} else if err != nil && maxTries==0 {
			fmt.Println(err)
			panic(err)
		}
		time.Sleep(delay * time.Millisecond)

	}
	sqlCommands := parse_sql.SQLCmds("setup.sql")
	for _, cmd := range sqlCommands {
		_, err = Db.Exec(cmd)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	post := Post{Content: "Hello db!", Author: "Jorge"}

	fmt.Println(post)
	post.Create()
	fmt.Println(post)

	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)

	readPost.Content = "Hoooolaaaa"
	readPost.Author = "michael weston"
	readPost.Update()

	posts, _ := GetPosts(5)
	fmt.Println(posts)

	readPost.Delete()
}
