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

type InMemoryTodoRepositoryImpl struct {
	todos models.Todos
}

func InitializeInMemoryTodoRepository() TodoRepository {

	defaultTodos := models.Todos{
		models.Todo{Id: uuid.MustParse("f6dd9451-ce63-40e6-af3c-66c4d5b4495d"), Name: "Yes", Completed: false, CreatedOn: time.Now()},
	}

	localStorageTodoRepo := InMemoryTodoRepositoryImpl{
		todos: defaultTodos,
	}
	return &localStorageTodoRepo
}

func (repo *InMemoryTodoRepositoryImpl) GetTodo(todoId uuid.UUID) (models.Todo, bool) {
	for _, v := range repo.todos {
		if v.Id == todoId {
			return v, true
		}
	}
	return models.Todo{}, false
}

func (repo *InMemoryTodoRepositoryImpl) GetAllTodos() models.Todos {
	return repo.todos
}

func (repo *InMemoryTodoRepositoryImpl) EditTodo(todo models.Todo) (models.Todo, error) {
	for index, v := range repo.todos {
		if v.Id == todo.Id {
			repo.todos[index] = todo
			return todo, nil
		}
	}
	return models.Todo{}, errors.New("todo to be edited not found , make sure AddTodo was called before edit")
}

func (repo *InMemoryTodoRepositoryImpl) AddTodo(todo models.Todo) models.Todo {
	//prepend the todos at top of slice
	repo.todos = append([]models.Todo{todo}, repo.todos...)
	return todo
}
