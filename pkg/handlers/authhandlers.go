package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/mananwalia959/go-todos-app/pkg/config"
	"github.com/mananwalia959/go-todos-app/pkg/models"
)

func InitializeAuthHandlers(appconfig config.Appconfig) AuthHandler {
	return AuthHandler{clientId: appconfig.OauthClientId, clientSecret: appconfig.OauthClientSecret}
}

type AuthHandler struct {
	clientId     string
	clientSecret string
}

func (authHandler AuthHandler) GetLoginUrl(w http.ResponseWriter, r *http.Request) {
	link := "https://accounts.google.com/o/oauth2/v2/auth"
	parsed, _ := url.Parse(link)
	q := parsed.Query()

	q.Add("client_id", authHandler.clientId)
	q.Add("redirect_uri", "http://localhost:3000/callback/googleoauth")
	q.Add("response_type", "code")
	q.Add("scope", "https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile")
	q.Add("state", "randomstring") //will be made random later

	parsed.RawQuery = q.Encode()

	http.Redirect(w, r, parsed.String(), http.StatusTemporaryRedirect)

}

type accessTokenReq struct {
	Client_id     string `json:"client_id"`
	Client_secret string `json:"client_secret"`
	Grant_type    string `json:"grant_type"`
	Redirect_uri  string `json:"redirect_uri"`
	AuthCode      string `json:"code"`
}

func (authHandler AuthHandler) GetToken(w http.ResponseWriter, r *http.Request) {
	tokenRequest := models.TokenRequest{}
	err := json.NewDecoder(r.Body).Decode(&tokenRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	url := "https://oauth2.googleapis.com/token"
	method := "POST"

	Code := tokenRequest.Code

	data := struct {
		Client_id     string `json:"client_id"`
		Client_secret string `json:"client_secret"`
		Grant_type    string `json:"grant_type"`
		Redirect_uri  string `json:"redirect_uri"`
		AuthCode      string `json:"code"`
	}{
		Client_id:     authHandler.clientId,
		Client_secret: authHandler.clientSecret,
		Grant_type:    "authorization_code",
		Redirect_uri:  "http://localhost:3000/callback/googleoauth",
		AuthCode:      Code,
	}

	payload, _ := json.Marshal(data)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(payload))

	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

}
