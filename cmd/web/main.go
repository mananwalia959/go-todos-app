package main

import (
	"fmt"
	"net/http"

	app "github.com/mananwalia959/go-todos-app/pkg/app"
)

var port = ":8080"

func main() {
	myRouter := app.GetApplication()

	fmt.Println("Starting port on ", port)
	http.ListenAndServe(port, myRouter)

}
