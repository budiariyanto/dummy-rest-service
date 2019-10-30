package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/budiariyanto/dummy-rest-service/model"
	"github.com/go-chi/chi"
)

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := todoService.GetAllTodos()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	wrapper := map[string]interface{}{
		"data": todos,
	}

	todosBody, err := json.Marshal(wrapper)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(todosBody)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	iID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	todos, err := todoService.GetTodo(iID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	todosBody, err := json.Marshal(todos)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(todosBody)
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	dcoder := json.NewDecoder(r.Body)

	err := dcoder.Decode(&todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = todoService.Insert(&todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(`{"status": "OK"}`))
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	dcoder := json.NewDecoder(r.Body)

	err := dcoder.Decode(&todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = todoService.Update(&todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(`{"status": "OK"}`))
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	dcoder := json.NewDecoder(r.Body)

	err := dcoder.Decode(&todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = todoService.Delete(&todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(`{"status": "OK"}`))
}
