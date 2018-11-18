package main

import "time"

// Message is a structure of messages to be posted to guest book.
type Message struct {
	Name      string    `datastore:"name"`
	Message   string    `datastore:"message"`
	Created   string    `datastore:"created"`
	CreatedAt time.Time `datastore:"createdAt"`
}
