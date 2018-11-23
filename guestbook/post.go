package main

import (
	"net/http"
	"strconv"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func post(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	ctx, err := appengine.Namespace(c, r.Host)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
		CreatedAt: time.Now(),
	}

	var key *datastore.Key
	k := r.FormValue("key")

	if k == "" {
		key = datastore.NewIncompleteKey(ctx, "Message", nil)
	} else {
		keyID, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		key = datastore.NewKey(ctx, "Message", "", keyID, nil)
	}

	if _, err := datastore.Put(ctx, key, msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
