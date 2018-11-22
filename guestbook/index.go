package main

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

var indexTmpl = template.Must(template.ParseFiles("./view/index.html"))

var title = "Wedding Guest Book"
var description = "Gopher & Gopher 2018.11.25"

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
	q := datastore.NewQuery("Message").Order("-createdAt").Limit(15)
	if _, err := q.GetAll(ctx, &msgs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
