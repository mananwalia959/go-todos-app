package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mananwalia959/go-todos-app/pkg/config"
)

func GetApplication(appconfig config.Appconfig) http.Handler {

	myRouter := mux.NewRouter()

	apiRoutes := myRouter.PathPrefix("/api").Subrouter()
	setApiRoutes(apiRoutes, appconfig)

	myRouter.PathPrefix("/").Handler(spaHandler())

	return myRouter

}

func spaHandler() http.Handler {
	return http.FileServer(http.Dir("client/build"))
}
