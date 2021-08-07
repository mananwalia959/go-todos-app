package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mananwalia959/go-todos-app/pkg/models"
	"github.com/mananwalia959/go-todos-app/pkg/repository"
	"github.com/mananwalia959/go-todos-app/pkg/utils"
)

func InitialzeTodoHandlers(todorepo repository.TodoRepository) TodosHandler {
	return TodosHandler{todoRepository: todorepo}
}

type TodosHandler struct {
	todoRepository repository.TodoRepository
}

func (handler TodosHandler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	allTodos := handler.todoRepository.GetAllTodos(r.Context())
	encodeToJson(w, 200, allTodos)
}

func (handler TodosHandler) GetSingleTodo(w http.ResponseWriter, r *http.Request) {
	todoId, err := getTodoId(r)
	if err != nil {
		ErrorResponse(w, 400, "please enter valid todoId")
		return
	}

	todo, found := handler.todoRepository.GetTodo(r.Context(), todoId)
	if !found {
		ErrorResponse(w, 404, "todo not found")
		return
	}
	encodeToJson(w, 200, todo)

}

func (handler TodosHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoId, err := getTodoId(r)
	if err != nil {
		ErrorResponse(w, 400, "please enter valid todoId")
		return
	}
	todo, found := handler.todoRepository.GetTodo(r.Context(), todoId)
	if !found {
		ErrorResponse(w, 404, "todo not found")
		return
	}
	result := handler.todoRepository.DeleteTodo(r.Context(), todo)
	if !result {
		ErrorResponse(w, 404, "Something went wrong")
		return
	}
	//empty response
	resp := struct{}{}
	encodeToJson(w, 204, resp)
}

func (handler TodosHandler) EditTodo(w http.ResponseWriter, r *http.Request) {
	todoId, err := getTodoId(r)
	if err != nil {
		ErrorResponse(w, 400, "please enter valid todoId")
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

	todo, found := handler.todoRepository.GetTodo(r.Context(), todoId)
	if !found {
		ErrorResponse(w, 404, "todo not found")
		return
	}

	todo.Name = editRequest.Name
	todo.Description = editRequest.Description
	todo.Completed = editRequest.Completed

	todo, successful := handler.todoRepository.EditTodo(r.Context(), todo)
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
		CreatedBy:   utils.GetUserPrincipal(r).Id,
	}
	savedTodo := handler.todoRepository.AddTodo(r.Context(), todo)
	encodeToJson(w, 200, savedTodo)
}

func getTodoId(r *http.Request) (uuid.UUID, error) {
	vars := mux.Vars(r)
	todoString := vars["todoId"]
	todoId, err := uuid.Parse(todoString)
	return todoId, err
}
