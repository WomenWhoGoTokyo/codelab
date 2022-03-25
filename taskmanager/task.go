package taskmanager

import (
	"context"
	"os"
	"time"

	"cloud.google.com/go/datastore"
)

// Task is the structure of the task.
type Task struct {
	ID        int64     `datastore:"-"`
	Title     string    `datastore:"title"`
	Status    Status    `datastore:"status"`
	CreatedAt time.Time `datastore:"createdAt"`
}

func newTask(title string) *Task {
	return &Task{
		Title:     title,
		Status:    ToDo,
		CreatedAt: time.Now(),
	}
}

func (t Task) add() error {
	ctx := context.Background()

	client, err := datastore.NewClient(ctx, os.Getenv("GCP_PROJECT"))
	if err != nil {
		return err
	}

	newKey := datastore.IncompleteKey(os.Getenv("MY_CODE"), nil)
	_, err = client.Put(ctx, newKey, &t)
	return err
}
