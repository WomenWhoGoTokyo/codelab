package main

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
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

func index(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	ctx, err := appengine.Namespace(c, r.Host)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var msgs []*Message
	q := datastore.NewQuery("Message").Order("-createdAt").Limit(10)
	keys, err := q.GetAll(ctx, &msgs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for i := 0; i < len(msgs); i++ {
		msgs[i].ID = keys[i].IntID()
	}

	idxt := &IndexTemplate{
		Title:       title,
		Description: description,
		Count:       len(msgs),
		Messages:    msgs,
	}

	if err := indexTmpl.Execute(w, idxt); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
