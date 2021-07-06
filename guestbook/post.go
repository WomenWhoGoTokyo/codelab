package guestbook

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"context"

	"cloud.google.com/go/datastore"
)

func Post(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, "wwgt-codelabs")
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
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
		key = datastore.IncompleteKey(r.Host, nil)
	} else {
		keyID, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusBadRequest)
			return
		}
		key = datastore.IDKey(r.Host, keyID, nil)
	}

	if _, err := client.Put(ctx, key, msg); err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
