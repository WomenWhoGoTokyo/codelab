package main

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

var indexTmpl = template.Must(template.ParseFiles("./view/index.html"))

func index(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	var msgs []*Message
	q := datastore.NewQuery("Message").Order("-createdAt").Limit(10)
	if _, err := q.GetAll(ctx, &msgs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := indexTmpl.Execute(w, msgs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
