package taskmanager

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

// TaskManager is entry point of Cloud Functions.
func TaskManager(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// 新規作成
	case http.MethodPost:
		param, code, err := getJSON(r.Header.Get("Content-Type"), r.Body)
		if err != nil {
			responseWrite(w, code, err.Error(), err)
			return
		}

		t := newTask(param.Title)
		if err := t.add(); err != nil {
			e := errors.Wrap(err, "post error")
			responseWrite(w, http.StatusInternalServerError, e.Error(), e)
			return
		}
		msg := fmt.Sprintf("%v added", t.Title)
		responseWrite(w, http.StatusOK, msg, nil)

	// 一覧取得
	case http.MethodGet:

	// ステータス変更
	case http.MethodPatch:

	// 削除
	case http.MethodDelete:

	default:
		e := errors.New("method not allowed")
		responseWrite(w, http.StatusMethodNotAllowed, e.Error(), e)
	}

	return
}

func responseWrite(w http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(code)
	w.Write([]byte(msg))
}

func getJSON(contentType string, body io.Reader) (Parameter, int, error) {
	var p Parameter

	if contentType != "application/json" {
		cause := errors.New("contentType")
		e := errors.Wrap(cause, "bad request Content-Type")
		return p, http.StatusBadRequest, e
	}

	b, err := ioutil.ReadAll(body)
	if err != nil {
		e := errors.Wrap(err, "read error")
		return p, http.StatusInternalServerError, e
	}

	if err = json.Unmarshal(b, &p); err != nil {
		e := errors.Wrap(err, "json parse error")
		return p, http.StatusInternalServerError, e
	}

	return p, http.StatusOK, nil
}
