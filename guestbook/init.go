package main

import "net/http"

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/post", post)
	http.HandleFunc("/edit", edit)
}
