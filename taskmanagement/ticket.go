package taskmanagement

import (
	"context"
	"os"
	"time"

	"cloud.google.com/go/datastore"
)

// Ticket is the structure of the task.
type Ticket struct {
	ID        int64     `datastore:"-"`
	Title     string    `datastore:"title"`
	Status    Status    `datastore:"status"`
	CreatedAt time.Time `datastore:"createdAt"`
}

func newTicket(title string) *Ticket {
	return &Ticket{
		Title:     title,
		Status:    ToDo,
		CreatedAt: time.Now(),
	}
}

func (t Ticket) add() error {
	ctx := context.Background()

	client, err := datastore.NewClient(ctx, "wwgt-codelabs")
	if err != nil {
		return err
	}

	newKey := datastore.IncompleteKey(os.Getenv("MY_CODE"), nil)
	if _, err := client.Put(ctx, newKey, &t); err != nil {
		return err
	}
	return nil
}
