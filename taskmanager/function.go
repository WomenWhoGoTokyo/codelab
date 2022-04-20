package taskmanager

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
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
			responseWrite(w, http.StatusInternalServerError, err.Error(), err)
			return
		}
		msg := fmt.Sprintf("%v added", t.Title)
		responseWrite(w, http.StatusOK, msg, nil)

	// 一覧取得
	case http.MethodGet:
		t, err := getAllTask()
		if err != nil {
			e := errors.Errorf("get error: %v", err)
			responseWrite(w, http.StatusInternalServerError, e.Error(), e)
			return
		}

		if len(t) < 1 {
			responseWrite(w, http.StatusOK, "0 tickets", nil)
			return
		}

		b, err := json.Marshal(t)
		if err != nil {
			e := errors.Errorf("json marshal error: %v", err)
			responseWrite(w, http.StatusInternalServerError, e.Error(), e)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(b)

	// ステータス変更
	case http.MethodPatch:
		param, code, err := getJSON(r.Header.Get("Content-Type"), r.Body)
		if err != nil {
			responseWrite(w, code, err.Error(), err)
			return
		}

		t := setTask(param.ID, param.Title, param.Status)
		if err := t.update(); err != nil {
			e := errors.Errorf("patch error: %v", err)
			responseWrite(w, http.StatusInternalServerError, e.Error(), e)
			return
		}
		msg := fmt.Sprintf("%v updated", t)
		responseWrite(w, http.StatusOK, msg, nil)

	// 削除
	case http.MethodDelete:
		param, code, err := getJSON(r.Header.Get("Content-Type"), r.Body)
		if err != nil {
			responseWrite(w, code, err.Error(), err)
			return
		}

		t := setTask(param.ID, "", 0)
		if err := t.delete(); err != nil {
			e := errors.Errorf("delete error: %v", err)
			responseWrite(w, http.StatusInternalServerError, e.Error(), e)
			return
		}
		msg := fmt.Sprintf("%v deleted", t)
		responseWrite(w, http.StatusOK, msg, nil)

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
		return p, http.StatusBadRequest, errors.New("invalid Content Type")
	}

	b, err := ioutil.ReadAll(body)
	if err != nil {
		return p, http.StatusInternalServerError, err
	}

	if err = json.Unmarshal(b, &p); err != nil {
		return p, http.StatusInternalServerError, err
	}

	return p, http.StatusOK, nil
}
