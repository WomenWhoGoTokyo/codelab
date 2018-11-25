package main

import (
	"net/http"

	"github.com/sinmetal/codelab/guestbook"
	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/post", guestbook.Post)
	http.HandleFunc("/", guestbook.Index)

	appengine.Main()
}
