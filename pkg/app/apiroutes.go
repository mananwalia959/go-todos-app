package app

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mananwalia959/go-todos-app/pkg/config"
	"github.com/mananwalia959/go-todos-app/pkg/handlers"
	"github.com/mananwalia959/go-todos-app/pkg/handlers/middlewares"
	"github.com/mananwalia959/go-todos-app/pkg/oauth"
	"github.com/mananwalia959/go-todos-app/pkg/repository"
)

func setApiRoutes(apiRoutes *mux.Router, appconfig config.Appconfig) {

	apiRoutes.Use(middlewares.PanicRecovermiddleWare)

	pool := repository.GetPool(&appconfig)

	todoRepo := repository.InitializeInMemoryTodoRepository()
	userRepo := repository.InitializePostgresUserRepository(pool)

	jwtUtil := oauth.InitializeJwtUtil(&appconfig)

	todoHandler := handlers.InitialzeTodoHandlers(todoRepo)
	authHandler := handlers.InitializeAuthHandlers(&appconfig, userRepo, jwtUtil)

	unProtectedRoutes := apiRoutes.NewRoute().Subrouter()

	unProtectedRoutes.Methods(http.MethodGet).Path("/auth/loginurl/google").HandlerFunc(authHandler.GetLoginUrl)
	unProtectedRoutes.Methods(http.MethodPost).Path("/auth/token/google").HandlerFunc(authHandler.GetToken)

	protectedRoutes := apiRoutes.NewRoute().Subrouter()
	protectedRoutes.Use(middlewares.GetAuthMiddleWare(jwtUtil))

	protectedRoutes.Methods(http.MethodGet).Path("/todos").HandlerFunc(todoHandler.GetAllTodos)
	protectedRoutes.Methods(http.MethodPost).Path("/todos").HandlerFunc(todoHandler.CreateTodo)
	protectedRoutes.Methods(http.MethodGet).Path("/todos/{todoId}").HandlerFunc(todoHandler.GetSingleTodo)
	protectedRoutes.Methods(http.MethodPut).Path("/todos/{todoId}").HandlerFunc(todoHandler.EditTodo)
	protectedRoutes.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)

}
