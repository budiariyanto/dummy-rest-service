package repository

import (
	"errors"
	"fmt"

	"github.com/budiariyanto/dummy-rest-service/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", "root", "12345", "localhost", 3306, "dummy-service", "charset=utf8&parseTime=True&loc=Asia%2fJakarta&time_zone=%27%2B07%3A00%27"))
	if err != nil {
		panic(err)
	}
}

type TodoRepo struct{}

type ITodo interface {
	GetAllTodos() (todos []model.Todo, err error)
	GetTodo(id int64) (todo model.Todo, err error)
	Insert(todo *model.Todo) (err error)
	Update(todo *model.Todo) (err error)
	Delete(todo *model.Todo) (err error)
}

func NewTodoRepo() *TodoRepo {
	return &TodoRepo{}
}

func (t *TodoRepo) GetAllTodos() (todos []model.Todo, err error) {
	err = db.Find(&todos).Error
	return
}

func (t *TodoRepo) GetTodo(id int64) (todo model.Todo, err error) {
	err = db.First(&todo, id).Error
	return
}

func (t *TodoRepo) Insert(todo *model.Todo) (err error) {
	return db.Create(todo).Error
}

func (t *TodoRepo) Update(todo *model.Todo) (err error) {
	return db.Save(todo).Error
}

func (t *TodoRepo) Delete(todo *model.Todo) (err error) {
	if todo.ID == 0 {
		return errors.New("primary key cannot be null")
	}
	return db.Delete(todo).Error
}
