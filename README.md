# go-todos-app

## Description
Simple Web App for Golang and React

## Pipelines

|Name| Pipeline|
| --- | --- |
|Build|[![go-todos-app](https://github.com/mananwalia959/go-todos-app/actions/workflows/pipeline.yml/badge.svg)](https://github.com/mananwalia959/go-todos-app/actions/workflows/pipeline.yml)|


## Required Env variables
|ENV VARIABLE NAME | WHERE TO GET IT | PURPOSE |
|---| --- | --- |
| OAUTH_CLIENT_ID_GOOGLE | your google oauth panel , refer [here](https://developers.google.com/identity/protocols/oauth2/web-server)  | For Our Google Sign in Functionality |
| OAUTH_CLIENT_SECRET_GOOGLE | your google oauth panel , refer [here](https://developers.google.com/identity/protocols/oauth2/web-server)  | For Our Google Sign in Functionality |
| REDIRECT_URL | your callback url , just use your ui url + '/callback/googleoauth', for ex: http://localhost:3000/callback/googleoauth  | For Our Google Sign in Functionality |
| SECRET_KEY_JWT | a random string (preferably long and diificult to guess)  | For Signing and verifying our jwt tokens |




