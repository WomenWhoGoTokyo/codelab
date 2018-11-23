package main

import (
	"net/http"
)

// var indexTmpl = template.Must(template.ParseFiles("./view/edit.html"))

// var title = "Wedding Guest Book 編集"

/*
// EditTemplate is a structure of edit template.
type EditTemplate struct {
	Title   string
	Name    string
	Message string
	ID      string
}
*/
func edit(w http.ResponseWriter, r *http.Request) {
	/*
		c := appengine.NewContext(r)
		ctx, err := appengine.Namespace(c, r.Host)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		p := r.FormValue("m")
		if p == "" {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// TODO: 取得

		title := "ほげほげ"
		name := "なまえ"
		message := "メッセージ"
		id := "ID"

		edt := &EditTemplate{
			Title:   title,
			Name:    name,
			Message: message,
			ID:      id,
		}

		if err := indexTmpl.Execute(w, edt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	*/
}
