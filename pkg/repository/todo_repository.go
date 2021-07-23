package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mananwalia959/go-todos-app/pkg/models"
)

type TodoRepository interface {
	GetAllTodos(ctx context.Context) models.Todos
	GetTodo(ctx context.Context, todoId uuid.UUID) (models.Todo, bool)
	AddTodo(ctx context.Context, todo models.Todo) models.Todo
	EditTodo(ctx context.Context, todo models.Todo) (models.Todo, bool)
}
