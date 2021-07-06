package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/WomenWhoGoTokyo/codelab/guestbook"
)

func main() {
	http.HandleFunc("/post", guestbook.Post)
	http.HandleFunc("/", guestbook.Index)
	http.HandleFunc("/edit", guestbook.Edit)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
