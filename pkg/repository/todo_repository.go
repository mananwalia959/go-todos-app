package repository

import (
	"github.com/google/uuid"
	"github.com/mananwalia959/go-todos-app/pkg/models"
)

type TodoRepository interface {
	GetAllTodos() models.Todos
	GetTodo(todoId uuid.UUID) (models.Todo, bool)
}

type localStorageTodoRepository struct {
}

func GetTodoRepository() TodoRepository {
	localStorageTodoRepo := localStorageTodoRepository{}
	return &localStorageTodoRepo
}

var todos = models.Todos{
	models.Todo{Id: uuid.MustParse("f6dd9451-ce63-40e6-af3c-66c4d5b4495d"), Name: "Yes", Completed: false},
}

func (repo *localStorageTodoRepository) GetTodo(todoId uuid.UUID) (models.Todo, bool) {
	for _, v := range todos {
		if v.Id == todoId {
			return v, true
		}
	}
	return models.Todo{}, false
}

func (repo *localStorageTodoRepository) GetAllTodos() models.Todos {
	return todos
}
