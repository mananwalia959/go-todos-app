package main

import (
	"fmt"
	"log"
	"net/http"

	app "github.com/mananwalia959/go-todos-app/pkg/app"
)

var port = ":8080"

func main() {
	myRouter := app.GetApplication()

	fmt.Println("Starting port on ", port)
	err := http.ListenAndServe(port, myRouter)
	// ListenAndServe will block the main goRoutine untill there is an error
	log.Fatal(err)

}
