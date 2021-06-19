package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mananwalia959/go-todos-app/pkg/config"
	"github.com/mananwalia959/go-todos-app/pkg/repository"
)

func GetApplication(appconfig config.Appconfig) http.Handler {

	todoRepo := repository.GetTodoRepository()

	myRouter := mux.NewRouter()

	apiRoutes := myRouter.PathPrefix("/api").Subrouter()
	setApiRoutes(apiRoutes, todoRepo, appconfig)

	myRouter.PathPrefix("/").Handler(spaHandler())

	return myRouter

}

func spaHandler() http.Handler {
	return http.FileServer(http.Dir("client/build"))
}
