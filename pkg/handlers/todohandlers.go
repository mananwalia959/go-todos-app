package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mananwalia959/go-todos-app/pkg/models"
	"github.com/mananwalia959/go-todos-app/pkg/repository"
)

func InitialzeTodoHandlers(todorepo repository.TodoRepository) TodosHandler {
	return TodosHandler{todoRepository: todorepo}
}

type TodosHandler struct {
	todoRepository repository.TodoRepository
}

func (handler TodosHandler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	allTodos := handler.todoRepository.GetAllTodos(context.TODO())
	encodeToJson(w, 200, allTodos)
}

func (handler TodosHandler) GetSingleTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	todoString := vars["todoId"]

	todoId, err := uuid.Parse(todoString)
	if err != nil {
		ErrorResponse(w, 400, "please enter valid uuid")
		return
	}

	todo, found := handler.todoRepository.GetTodo(context.TODO(), todoId)
	if !found {
		ErrorResponse(w, 404, "todo not found")
		return
	}
	encodeToJson(w, 200, todo)

}

func (handler TodosHandler) EditTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	todoString := vars["todoId"]

	todoId, err := uuid.Parse(todoString)
	if err != nil {
		ErrorResponse(w, 400, "please enter valid uuid")
		return
	}

	var editRequest models.TodoEditRequest

	err = json.NewDecoder(r.Body).Decode(&editRequest)
	if err != nil {
		ErrorResponse(w, 400, "Provide valid edit request")
		return
	}
	if len(strings.TrimSpace(editRequest.Name)) == 0 {
		ErrorResponse(w, 400, "name must not be empty")
		return
	}

	todo, found := handler.todoRepository.GetTodo(context.TODO(), todoId)
	if !found {
		ErrorResponse(w, 404, "todo not found")
		return
	}

	todo.Name = editRequest.Name
	todo.Description = editRequest.Description
	todo.Completed = editRequest.Completed

	todo, successful := handler.todoRepository.EditTodo(context.TODO(), todo)
	if !successful {
		ErrorResponse(w, 500, "Something went wrong")
		return
	}

	encodeToJson(w, 200, todo)

}

func (handler TodosHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var createRequest models.TodoCreateRequest
	err := json.NewDecoder(r.Body).Decode(&createRequest)
	if err != nil {
		ErrorResponse(w, 400, "Provide valid create request")
		return
	}
	if len(strings.TrimSpace(createRequest.Name)) == 0 {
		ErrorResponse(w, 400, "name must not be empty")
		return
	}
	todo := models.Todo{
		Id:          uuid.New(),
		Name:        createRequest.Name,
		Description: createRequest.Description,
		Completed:   false,
		CreatedOn:   time.Now(),
		CreatedBy:   GetUserPrincipal(r).Id,
	}
	savedTodo := handler.todoRepository.AddTodo(context.TODO(), todo)
	encodeToJson(w, 200, savedTodo)
}
