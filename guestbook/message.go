package guestbook

import "time"

// Message is a structure of messages to be posted to guest book.
type Message struct {
	ID        int64     `datastore:"-"`
	Name      string    `datastore:"name"`
	Message   string    `datastore:"message"`
	CreatedAt time.Time `datastore:"createdAt"`
}
