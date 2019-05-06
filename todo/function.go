package todo

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"cloud.google.com/go/datastore"
)

// Status is the progress of the task.
type Status int

const (
	// ToDo is the task progress status.
	ToDo Status = iota + 1
	// InProgress is the task progress status.
	InProgress
	// Done is the task progress status.
	Done
)

// Parameter is the parameter from webhook.
type Parameter struct {
	SubCommand string
	Task       Task
}

// Task is the structure of the task.
type Task struct {
	ID        int64     `datastore:"-"`
	Title     string    `datastore:"title"`
	Status    Status    `datastore:"status"`
	CreatedAt time.Time `datastore:"createdAt"`
}

func ToDo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		e := "Method Not Allowed."
		log.Println(e)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(e))
		return
	}

	// TODO: JSON で受け取る
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("ReadAllError: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	parsed, err := url.ParseQuery(string(b))
	if err != nil {
		log.Printf("ParseQueryError: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	p := new(Parameter)
	p.parse(parsed.Get("text"))

	switch p.SubCommand {
	case "add":
		if err := add(p.Value); err != nil {
			log.Printf("DatastorePutError: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(p.Value))
	case "list":
		list, err := list()
		if err != nil {
			log.Printf("DatastoreGetAllError: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(sprint(list)))
	default:
		e := "Invalid SubCommand."
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(e))
	}
	return
}

func (p *Parameter) parse(text string) {
	t := strings.TrimSpace(text)
	if len(t) < 1 {
		return
	}
	s := strings.SplitN(t, " ", 2)
	p.SubCommand = s[0]

	if len(s) == 1 {
		return
	}
	p.Value = s[1]
}

func add(value string) error {
	// TODO: 名前空間を指定する
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, "")
	if err != nil {
		return err
	}

	newKey := datastore.IncompleteKey("Task", nil)
	r := Task{
		Title:     value,
		Status:    ToDo,
		CreatedAt: time.Now(),
	}
	if _, err := client.Put(ctx, newKey, &r); err != nil {
		return err
	}
	return nil
}

func list() ([]Restaurant, error) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, "")
	if err != nil {
		return nil, err
	}

	var t []Task
	q := datastore.NewQuery("Task").Order("-createdAt").Limit(5)
	keys, err := client.GetAll(ctx, q, &t)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(t); i++ {
		t[i].ID = keys[i].ID
	}
	return t, nil
}

func sprint(list []Task) (s string) {
	for _, t := range list {
		s = s + fmt.Sprintf("[%v] %v\n", t.ID, t.Name)
	}
	return s
}
