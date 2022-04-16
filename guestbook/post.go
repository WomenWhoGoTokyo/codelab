package guestbook

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"context"

	"cloud.google.com/go/datastore"
)

func Post(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, os.Getenv("GOOGLE_CLOUD_PROJECT"))
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

	key := datastore.IncompleteKey(r.Host, nil)
	if _, err := client.Put(ctx, key, msg); err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
