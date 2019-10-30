package main

import (
	"github.com/budiariyanto/dummy-rest-service/handler"
	"github.com/go-chi/chi"
)

func SetRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/all", handler.GetAllTodos)
	r.Get("/todo/{id}", handler.GetTodo)
	r.Post("/add", handler.AddTodo)
	r.Put("/update", handler.UpdateTodo)
	r.Delete("/delete", handler.DeleteTodo)

	return r
}
