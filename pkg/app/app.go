package app

import (
	"net/http"
	"os"

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

/**
* inspired from https://github.com/gorilla/mux#serving-single-page-applications
* or https://github.com/gorilla/mux/pull/493/files
**/
func spaHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// prepend the path with the path to the static directory
		path = "./client/build/" + path

		// check whether a file exists at the given path
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			// file does not exist, serve index.html
			http.ServeFile(w, r, "./client/build/index.html")
			return
		} else if err != nil {
			// if we got an error (that wasn't that the file doesn't exist) stating the
			// file, return a 500 internal server error and stop
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// otherwise, use http.FileServer to serve the static dir
		http.FileServer(http.Dir("./client/build")).ServeHTTP(w, r)
	})
}
