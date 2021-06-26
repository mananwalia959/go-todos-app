package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/mananwalia959/go-todos-app/pkg/config"
	"github.com/mananwalia959/go-todos-app/pkg/models"
)

func InitializeAuthHandlers(appconfig config.Appconfig) AuthHandler {
	return AuthHandler{clientId: appconfig.OauthClientId, clientSecret: appconfig.OauthClientSecret, client: &http.Client{}}
}

type AuthHandler struct {
	clientId     string
	clientSecret string
	client       *http.Client
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

func (authHandler AuthHandler) GetToken(w http.ResponseWriter, r *http.Request) {
	tokenRequest := models.TokenRequest{}
	err := json.NewDecoder(r.Body).Decode(&tokenRequest)
	Code := tokenRequest.Code

	if err != nil {
		errorResponse(w, 400, "please provide valid token request")
		return
	}

	accessTokenRequest := models.AccessTokenReqGoogle{
		Client_id:     authHandler.clientId,
		Client_secret: authHandler.clientSecret,
		Grant_type:    "authorization_code",
		Redirect_uri:  "http://localhost:3000/callback/googleoauth",
		AuthCode:      Code,
	}
	accessTokenReponse, err := getAccessTokenFromCode(authHandler.client, accessTokenRequest)

	if err != nil {
		errorResponse(w, 500, "something went wrong")
		return
	}

	profile, err := getProfileFromOauthApi(accessTokenReponse.AccessToken, authHandler.client)
	if err != nil {
		errorResponse(w, 500, "something went wrong")
		return
	}

	encodeToJson(w, 200, profile)
}

func getProfileFromOauthApi(accessToken string, client *http.Client) (models.GoogleProfileInfo, error) {
	url := "https://www.googleapis.com/oauth2/v1/userinfo"

	method := http.MethodGet

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return models.GoogleProfileInfo{}, err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)

	res, err := client.Do(req)
	if err != nil {
		return models.GoogleProfileInfo{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return models.GoogleProfileInfo{}, err
	}

	googleProfileInfo := models.GoogleProfileInfo{}
	err = json.NewDecoder(res.Body).Decode(&googleProfileInfo)
	if err != nil {
		return models.GoogleProfileInfo{}, err
	}

	return googleProfileInfo, nil
}

func getAccessTokenFromCode(client *http.Client, tokenReq models.AccessTokenReqGoogle) (models.AccessTokenRespGoogle, error) {

	url := "https://oauth2.googleapis.com/token"
	method := http.MethodPost

	payload, err := json.Marshal(tokenReq)
	if err != nil {
		return models.AccessTokenRespGoogle{}, err
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(payload))

	if err != nil {
		return models.AccessTokenRespGoogle{}, err
	}

	res, err := client.Do(req)
	if err != nil {
		return models.AccessTokenRespGoogle{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return models.AccessTokenRespGoogle{}, err
	}

	accessTokenReponse := models.AccessTokenRespGoogle{}
	err = json.NewDecoder(res.Body).Decode(&accessTokenReponse)
	if err != nil {
		return models.AccessTokenRespGoogle{}, err
	}
	return accessTokenReponse, nil

}
