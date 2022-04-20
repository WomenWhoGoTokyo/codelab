package taskmanager

import (
	"context"
	"os"
	"time"

	"cloud.google.com/go/datastore"
)

// Task is the structure of the task.
type Task struct {
	ID        int64     `datastore:"-" json:"ID"`
	Title     string    `datastore:"title" json:"title"`
	Status    Status    `datastore:"status" json:"status"`
	CreatedAt time.Time `datastore:"createdAt" json:"createdAt"`
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

func getAllTask() ([]Task, error) {
	ctx := context.Background()

	client, err := datastore.NewClient(ctx, "wwgt-codelabs")
	if err != nil {
		return nil, err
	}

	var t []Task

	q := datastore.NewQuery(os.Getenv("MY_CODE"))
	keys, err := client.GetAll(ctx, q, &t)
	if err != nil {
		return nil, err
	}

	for i, key := range keys {
		t[i].ID = key.ID
	}
	return t, nil
}

func setTask(id int64, title string, status Status) *Task {
	return &Task{
		ID:        id,
		Title:     title,
		Status:    status,
		CreatedAt: time.Now(),
	}
}

func (t Task) update() error {
	ctx := context.Background()

	client, err := datastore.NewClient(ctx, "wwgt-codelabs")
	if err != nil {
		return err
	}

	key := datastore.IDKey(os.Getenv("MY_CODE"), t.ID, nil)

	_, err = client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {
		var gt Task

		if err := tx.Get(key, &gt); err != nil {
			return err
		}

		if t.Title == "" {
			t.Title = gt.Title
		}

		if err := t.Status.validate(); err != nil {
			t.Status = gt.Status
		}

		t.CreatedAt = gt.CreatedAt

		_, err := tx.Put(key, &t)
		return err
	})
	return err
}

func (t Task) delete() error {
	ctx := context.Background()

	client, err := datastore.NewClient(ctx, "wwgt-codelabs")
	if err != nil {
		return err
	}

	return client.Delete(ctx, datastore.IDKey(os.Getenv("MY_CODE"), t.ID, nil))
}
