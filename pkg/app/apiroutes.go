package app

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mananwalia959/go-todos-app/pkg/config"
	"github.com/mananwalia959/go-todos-app/pkg/handlers"
	"github.com/mananwalia959/go-todos-app/pkg/repository"
)

func setApiRoutes(apiRoutes *mux.Router, todorepo repository.TodoRepository, appconfig config.Appconfig) {

	todoHandler := handlers.InitialzeTodoHandlers(todorepo)
	authHandler := handlers.InitializeAuthHandlers(appconfig)

	unProtectedRoutes := apiRoutes.NewRoute().Subrouter()

	unProtectedRoutes.Methods(http.MethodGet).Path("/auth/loginurl/google").HandlerFunc(authHandler.GetLoginUrl)
	unProtectedRoutes.Methods(http.MethodPost).Path("/auth/token/google").HandlerFunc(authHandler.GetToken)

	protectedRoutes := apiRoutes.NewRoute().Subrouter()

	protectedRoutes.Methods(http.MethodGet).Path("/todos").HandlerFunc(todoHandler.GetAllTodos)
	protectedRoutes.Methods(http.MethodPost).Path("/todos").HandlerFunc(todoHandler.CreateTodo)
	protectedRoutes.Methods(http.MethodGet).Path("/todos/{todoId}").HandlerFunc(todoHandler.GetSingleTodo)
	protectedRoutes.Methods(http.MethodPut).Path("/todos/{todoId}").HandlerFunc(todoHandler.EditTodo)
	protectedRoutes.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)

}
