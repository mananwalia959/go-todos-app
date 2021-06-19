package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	app "github.com/mananwalia959/go-todos-app/pkg/app"
	"github.com/mananwalia959/go-todos-app/pkg/config"
)

var port = ":8080"

func main() {

	config := getConfig()
	myRouter := app.GetApplication(config)

	fmt.Println("Starting port on ", port)
	err := http.ListenAndServe(port, myRouter)
	// ListenAndServe will block the main goRoutine untill there is an error
	log.Fatal(err)

}

func getConfig() config.Appconfig {
	clientID, clientidPresent := os.LookupEnv("OAUTH_CLIENT_ID_GOOGLE")
	if !clientidPresent {
		log.Fatal("exiting : client id is not present")
	}

	clientSecret, clientSecretPresent := os.LookupEnv("OAUTH_CLIENT_SECRET_GOOGLE")
	if !clientSecretPresent {
		log.Fatal("exiting : client secret is not present")
	}

	// 765874532491-6lr571h6anapg0uaufue3gvlpuvb6388.apps.googleusercontent.com
	return config.Appconfig{OauthClientId: clientID, OauthClientSecret: clientSecret}
}
