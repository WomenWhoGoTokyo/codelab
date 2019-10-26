package guestbook

import (
	"net/http"
	"time"
	"context"

	"cloud.google.com/go/datastore"
)

func Post(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, "wwgt-codelabs")
	if err != nil {
		http.Error(w, "クライアント作成に失敗しました", http.StatusInternalServerError)
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

	key := datastore.IncompleteKey("Message", nil)
	if _, err := client.Put(ctx, key, msg); err != nil {
		http.Error(w, "データの保存に失敗しました", http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
