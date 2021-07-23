package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/mananwalia959/go-todos-app/pkg/models"
)

type TodoRepository interface {
	GetAllTodos(ctx context.Context) models.Todos
	GetTodo(ctx context.Context, todoId uuid.UUID) (models.Todo, bool)
	AddTodo(ctx context.Context, todo models.Todo) models.Todo
	EditTodo(ctx context.Context, todo models.Todo) (models.Todo, bool)
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

func (repo *InMemoryTodoRepositoryImpl) GetTodo(ctx context.Context, todoId uuid.UUID) (models.Todo, bool) {
	for _, v := range repo.todos {
		if v.Id == todoId {
			return v, true
		}
	}
	return models.Todo{}, false
}

func (repo *InMemoryTodoRepositoryImpl) GetAllTodos(ctx context.Context) models.Todos {
	return repo.todos
}

/**
* Returns models.todo , true if successful
* empty models.todo , false if unsuccessful
*
 */
func (repo *InMemoryTodoRepositoryImpl) EditTodo(ctx context.Context, todo models.Todo) (models.Todo, bool) {
	for index, v := range repo.todos {
		if v.Id == todo.Id {
			repo.todos[index] = todo
			return todo, true
		}
	}
	return models.Todo{}, false
}

func (repo *InMemoryTodoRepositoryImpl) AddTodo(ctx context.Context, todo models.Todo) models.Todo {
	//prepend the todos at top of slice
	repo.todos = append([]models.Todo{todo}, repo.todos...)
	return todo
}
