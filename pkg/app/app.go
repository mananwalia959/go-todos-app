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

	apiRoutes := myRouter.PathPrefix("/api").Subrouter()

	myRouter.PathPrefix("/").Handler(spaHandler())

	setApiRoutes(apiRoutes)

	return myRouter

}

func spaHandler() http.Handler {
	return http.FileServer(http.Dir("client/build"))
}
