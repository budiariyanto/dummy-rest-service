package service

import (
	"github.com/budiariyanto/dummy-rest-service/model"
	"github.com/budiariyanto/dummy-rest-service/repository"
)

type IServiceTodo interface {
	GetAllTodos() (todos []model.Todo, err error)
	GetTodo() (todo model.Todo, err error)
	Insert(todo *model.Todo) (err error)
	Update(todo *model.Todo) (err error)
	Delete(todo *model.Todo) (err error)
}

type TodoService struct {
	repo repository.ITodo
}

func NewTodoService(repo repository.ITodo) *TodoService {
	return &TodoService{repo: repo}
}

func (t *TodoService) GetAllTodos() (todos []model.Todo, err error) {
	return t.repo.GetAllTodos()
}

func (t *TodoService) GetTodo(id int64) (todo model.Todo, err error) {
	return t.repo.GetTodo(id)
}

func (t *TodoService) Insert(todo *model.Todo) (err error) {
	return t.repo.Insert(todo)
}

func (t *TodoService) Update(todo *model.Todo) (err error) {
	return t.repo.Update(todo)
}

func (t *TodoService) Delete(todo *model.Todo) (err error) {
	return t.repo.Delete(todo)
}
