package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mananwalia959/go-todos-app/pkg/models"
	"github.com/mananwalia959/go-todos-app/pkg/utils"
)

func InitializePostgresTodoRepository(pool *pgxpool.Pool) TodoRepository {
	return &TodoRepositoryPostgresImpl{pool: pool}
}

type TodoRepositoryPostgresImpl struct {
	pool *pgxpool.Pool
}

func (repo *TodoRepositoryPostgresImpl) GetTodo(ctx context.Context, todoId uuid.UUID) (models.Todo, bool) {

	up := utils.GetUserPrincipalFromContext(ctx)
	query := "select id, name, description, completed, created_on, " +
		"created_by ,archived from todos where created_by = $1 AND id = $2"

	rows, err := repo.pool.Query(ctx, query, up.Id, todoId)
	panicIfNotNil(err)
	defer rows.Close()

	if rows.Next() {
		todo := readRow(rows)
		return todo, true
	}

	return models.Todo{}, false
}

func (repo *TodoRepositoryPostgresImpl) GetAllTodos(ctx context.Context) models.Todos {
	up := utils.GetUserPrincipalFromContext(ctx)
	query := "select id, name, description, completed, created_on, " +
		"created_by ,archived from todos where created_by = $1 order by created_on desc"

	rows, err := repo.pool.Query(ctx, query, up.Id)
	panicIfNotNil(err)
	defer rows.Close()
	//initialize to 8 by default (need more metrics to caliberate)
	todos := models.Todos{}
	for rows.Next() {
		todo := readRow(rows)
		todos = append(todos, todo)
	}
	return todos
}

func (repo *TodoRepositoryPostgresImpl) AddTodo(ctx context.Context, todo models.Todo) models.Todo {
	query := "insert into todos (id, name, description, completed, created_on ,created_by) values ($1,$2,$3,$4,$5,$6)"
	_, err := repo.pool.Exec(ctx, query, todo.Id, todo.Name, todo.Description, todo.Completed, todo.CreatedOn, todo.CreatedBy)
	panicIfNotNil(err)
	return todo
}

func (repo *TodoRepositoryPostgresImpl) EditTodo(ctx context.Context, todo models.Todo) (models.Todo, bool) {
	query := "update todos set name = $1, description = $2 ,completed = $3 where id = $4"

	res, err := repo.pool.Exec(ctx, query, todo.Name, todo.Description, todo.Completed, todo.Id)
	panicIfNotNil(err)
	return todo, res.RowsAffected() > 0
}

func panicIfNotNil(err error) {
	if err != nil {
		panic(err)
	}
}

// order is id, name, description, completed, created_on ,created_by ,archived
func readRow(rows pgx.Rows) models.Todo {
	var id uuid.UUID
	var name string
	var description string
	var completed bool
	var createdOn time.Time
	var created_by uuid.UUID
	var archived bool

	err := rows.Scan(&id, &name, &description, &completed, &createdOn, &created_by, &archived)
	panicIfNotNil(err)

	return models.Todo{
		Id:          id,
		Name:        name,
		Description: description,
		Completed:   completed,
		CreatedOn:   createdOn,
		CreatedBy:   created_by,
	}
}
