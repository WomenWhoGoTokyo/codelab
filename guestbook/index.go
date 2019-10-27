package guestbook

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"cloud.google.com/go/datastore"
)

var indexTmpl = template.Must(template.ParseFiles("./view/index.html"))

var title = "ゲストブック"
var description = "結婚式などの受付で名前や住所, メッセージを記帳してもらうためのノートのことです。"

// IndexTemplate is a structure of index template.
type IndexTemplate struct {
	Title       string
	Description string
	Count       int
	Messages    []*Message
}

func Index(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, "wwgt-codelabs")
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	var msgs []*Message
	q := datastore.NewQuery(r.Host).Order("-createdAt").Limit(10)
	keys, err := client.GetAll(ctx, q, &msgs)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	for i := 0; i < len(msgs); i++ {
		msgs[i].ID = keys[i].ID
	}

	idxt := &IndexTemplate{
		Title:       title,
		Description: description,
		Count:       len(msgs),
		Messages:    msgs,
	}

	if err := indexTmpl.Execute(w, idxt); err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
	}
}
