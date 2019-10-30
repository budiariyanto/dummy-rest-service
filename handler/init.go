package handler

import (
	"github.com/budiariyanto/dummy-rest-service/repository"
	"github.com/budiariyanto/dummy-rest-service/service"
)

var todoService *service.TodoService

func init() {
	repo := repository.NewTodoRepo()
	todoService = service.NewTodoService(repo)
}
