package main

import (
	"net/http"

	"github.com/WomenWhoGoTokyo/codelab/guestbook"
)

func main() {
	http.HandleFunc("/post", guestbook.Post)
	http.HandleFunc("/", guestbook.Index)
}
