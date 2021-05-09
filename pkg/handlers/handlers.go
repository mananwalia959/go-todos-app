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
)

var todoRepository repository.TodoRepository

func InitialzeHandlers(todorepo repository.TodoRepository) {
	todoRepository = todorepo
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	allTodos := todoRepository.GetAllTodos()
	encodeToJson(w, 200, allTodos)
}

func GetSingleTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	todoString := vars["todoId"]

	todoId, err := uuid.Parse(todoString)
	if err != nil {
		errorResponse(w, 400, "please enter valid uuid")
		return
	}

	todo, found := todoRepository.GetTodo(todoId)
	if !found {
		errorResponse(w, 404, "todo not found")
		return
	}
	encodeToJson(w, 200, todo)

}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var createRequest models.TodoCreateRequest
	err := json.NewDecoder(r.Body).Decode(&createRequest)
	if err != nil {
		errorResponse(w, 400, "Provide valid create request")
		return
	}
	if len(strings.TrimSpace(createRequest.Name)) == 0 {
		errorResponse(w, 400, "name must not be empty")
		return
	}
	todo := models.Todo{
		Id:          uuid.New(),
		Name:        createRequest.Name,
		Description: createRequest.Description,
		Completed:   false,
		CreatedOn:   time.Now(),
	}
	savedTodo := todoRepository.AddTodo(todo)
	encodeToJson(w, 200, savedTodo)
}

func encodeToJson(w http.ResponseWriter, status int, jsonInterface interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(jsonInterface)
}

func errorResponse(w http.ResponseWriter, status int, message string) {
	errorMessage := models.ErrorResponse{Message: message}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errorMessage)
}
