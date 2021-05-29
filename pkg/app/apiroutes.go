package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mananwalia959/go-todos-app/pkg/handlers"
)

func setApiRoutes(apiRoutes *mux.Router) {

	apiRoutes.Methods(http.MethodGet).Path("/todos").HandlerFunc(handlers.GetAllTodos)
	apiRoutes.Methods(http.MethodPost).Path("/todos").HandlerFunc(handlers.CreateTodo)
	apiRoutes.Methods(http.MethodGet).Path("/todos/{todoId}").HandlerFunc(handlers.GetSingleTodo)
	apiRoutes.Methods(http.MethodPut).Path("/todos/{todoId}").HandlerFunc(handlers.EditTodo)
	apiRoutes.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)

}
