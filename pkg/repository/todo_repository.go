package repository

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mananwalia959/go-todos-app/pkg/models"
)

type TodoRepository interface {
	GetAllTodos() models.Todos
	GetTodo(todoId uuid.UUID) (models.Todo, bool)
	AddTodo(todo models.Todo) models.Todo
	EditTodo(todo models.Todo) (models.Todo, error)
}

type localStorageTodoRepository struct {
	todos models.Todos
}

func GetTodoRepository() TodoRepository {

	defaultTodos := models.Todos{
		models.Todo{Id: uuid.MustParse("f6dd9451-ce63-40e6-af3c-66c4d5b4495d"), Name: "Yes", Completed: false, CreatedOn: time.Now()},
	}

	localStorageTodoRepo := localStorageTodoRepository{
		todos: defaultTodos,
	}
	return &localStorageTodoRepo
}

func (repo *localStorageTodoRepository) GetTodo(todoId uuid.UUID) (models.Todo, bool) {
	for _, v := range repo.todos {
		if v.Id == todoId {
			return v, true
		}
	}
	return models.Todo{}, false
}

func (repo *localStorageTodoRepository) GetAllTodos() models.Todos {
	return repo.todos
}

func (repo *localStorageTodoRepository) EditTodo(todo models.Todo) (models.Todo, error) {
	for index, v := range repo.todos {
		if v.Id == todo.Id {
			repo.todos[index] = todo
			return todo, nil
		}
	}
	return models.Todo{}, errors.New("Todo to be edited not found , make sure AddTodo was called before edit")
}

func (repo *localStorageTodoRepository) AddTodo(todo models.Todo) models.Todo {
	repo.todos = append(repo.todos, todo)
	return todo
}
