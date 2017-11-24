package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
)

type Post struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/post/", handleRequest)
	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		return
	}
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	var post Post
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		return
	}
	err = post.Create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id)
	if err != nil {
		return
	}
	err = json.NewDecoder(r.Body).Decode(post)
	if err != nil {
		return
	}
	err = post.Update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id)
	if err != nil {
		return
	}
	err = post.Delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
