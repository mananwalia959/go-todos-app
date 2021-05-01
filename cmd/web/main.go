package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mananwalia959/go-todos-app/pkg/models"
)

var port = ":8080"

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Methods(http.MethodGet).Path("/todos").HandlerFunc(todosHandler)

	fmt.Println("Starting port on ", port)
	http.ListenAndServe(port, myRouter)

}
func todosHandler(w http.ResponseWriter, r *http.Request) {

	todos := models.Todos{
		models.Todo{Id: uuid.MustParse("f6dd9451-ce63-40e6-af3c-66c4d5b4495d"), Name: "Yes", Completed: false},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
