package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/mananwalia959/go-todos-app/pkg/app"
	"github.com/mananwalia959/go-todos-app/pkg/config"
)

var port = ":8080"

func main() {

	config := getConfig()
	log.Println("Starting application")
	myRouter := app.GetApplication(config)

	log.Println("Starting on port", port)
	err := http.ListenAndServe(port, myRouter)
	// ListenAndServe will block the main goRoutine untill there is an error
	log.Fatal(err)

}

func getConfig() config.Appconfig {
	clientID := getEnv("OAUTH_CLIENT_ID_GOOGLE")
	clientSecret := getEnv("OAUTH_CLIENT_SECRET_GOOGLE")
	redirectUrl := getEnv("REDIRECT_URL")
	secretKeyJwt := getEnv("SECRET_KEY_JWT")
	postgresUrl := getEnv("POSTGRES_URL")
	postgresDbName := getEnv("POSTGRES_DB_NAME")
	postgresUserName := getEnv("POSTGRES_USERNAME")
	postgresPassword := getEnv("POSTGRES_PASSWORD")

	return config.Appconfig{
		OauthClientId:     clientID,
		OauthClientSecret: clientSecret,
		OauthRedirectUrl:  redirectUrl,
		SecretKeyJwt:      secretKeyJwt,
		PostgresUrl:       postgresUrl,
		PostgresUsername:  postgresUserName,
		PostgresDbName:    postgresDbName,
		PostgresPassword:  postgresPassword}
}

func getEnv(key string) string {
	value, found := os.LookupEnv(key)
	if !found {
		log.Fatal("exiting : ", key, " is not present")
	}
	return strings.TrimSpace(value)
}
