package main

import (
	"html/template"
	"net/http"
	"strconv"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

var editTmpl = template.Must(template.ParseFiles("./view/edit.html"))

// EditTemplate is a structure of edit template.
type EditTemplate struct {
	Title   string
	Name    string
	Message string
	KeyID   int64
}

func edit(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	ctx, err := appengine.Namespace(c, r.Host)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	k := r.FormValue("key")
	keyID, err := strconv.ParseInt(k, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var msg Message
	key := datastore.NewKey(ctx, "Message", "", keyID, nil)
	if err = datastore.Get(ctx, key, &msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	edt := &EditTemplate{
		Title:   title,
		Name:    msg.Name,
		Message: msg.Message,
		KeyID:   keyID,
	}

	if err := editTmpl.Execute(w, edt); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
