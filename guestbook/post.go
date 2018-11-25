package guestbook

import (
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func Post(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	ctx, err := appengine.Namespace(c, appengine.VersionID(c))
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

	key := datastore.NewIncompleteKey(ctx, "Message", nil)
	if _, err := datastore.Put(ctx, key, msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
