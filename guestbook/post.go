package main

import (
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func post(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	name := r.FormValue("name")
	if name == "" {
		name = "NO NAME"
	}

	message := r.FormValue("message")
	if message == "" {
		message = "NO MESSAGE"
	}

	msg := &Message{
		Name:      name,
		Message:   message,
		Created:   "",
		CreatedAt: time.Now(),
	}

	key := datastore.NewIncompleteKey(ctx, "Message", nil)
	if _, err := datastore.Put(ctx, key, msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
