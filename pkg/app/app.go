package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mananwalia959/go-todos-app/pkg/handlers"
	"github.com/mananwalia959/go-todos-app/pkg/repository"
)

func GetApplication() http.Handler {

	handlers.InitialzeHandlers(repository.GetTodoRepository())

	myRouter := mux.NewRouter()

	myRouter.Methods(http.MethodGet).Path("/todos").HandlerFunc(handlers.GetAllTodos)
	myRouter.Methods(http.MethodPost).Path("/todos").HandlerFunc(handlers.CreateTodo)
	myRouter.Methods(http.MethodGet).Path("/todos/{todoId}").HandlerFunc(handlers.GetSingleTodo)
	return myRouter

}
