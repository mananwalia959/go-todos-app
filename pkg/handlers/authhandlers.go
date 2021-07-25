package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/mananwalia959/go-todos-app/pkg/config"
	"github.com/mananwalia959/go-todos-app/pkg/models"
	"github.com/mananwalia959/go-todos-app/pkg/oauth"
	"github.com/mananwalia959/go-todos-app/pkg/repository"
)

func InitializeAuthHandlers(appconfig *config.Appconfig, userRepository repository.UserRepository, jwtutil oauth.JWTUtil) AuthHandler {
	return AuthHandler{
		clientId:       appconfig.OauthClientId,
		clientSecret:   appconfig.OauthClientSecret,
		client:         &http.Client{},
		redirectUrl:    appconfig.OauthRedirectUrl,
		userRepository: userRepository,
		jwtutil:        jwtutil,
	}
}

type AuthHandler struct {
	clientId       string
	clientSecret   string
	client         *http.Client
	redirectUrl    string
	userRepository repository.UserRepository
	jwtutil        oauth.JWTUtil
}

func (authHandler AuthHandler) GetLoginUrl(w http.ResponseWriter, r *http.Request) {
	link := "https://accounts.google.com/o/oauth2/v2/auth"
	parsed, _ := url.Parse(link)
	q := parsed.Query()

	q.Add("client_id", authHandler.clientId)
	q.Add("redirect_uri", authHandler.redirectUrl)
	q.Add("response_type", "code")
	q.Add("scope", "https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile")
	q.Add("state", "randomstring") //will be made random later

	parsed.RawQuery = q.Encode()

	http.Redirect(w, r, parsed.String(), http.StatusOK)

}

func (authHandler AuthHandler) GetToken(w http.ResponseWriter, r *http.Request) {
	tokenRequest := models.TokenRequest{}
	err := json.NewDecoder(r.Body).Decode(&tokenRequest)
	Code := tokenRequest.Code

	if err != nil {
		ErrorResponse(w, 400, "please provide valid token request")
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
		logIfErr(err)
		ErrorResponse(w, 500, "something went wrong")
		return
	}

	profile, err := getProfileFromOauthApi(accessTokenReponse.AccessToken, authHandler.client)
	if err != nil {
		logIfErr(err)
		ErrorResponse(w, 500, "something went wrong")
		return
	}

	userprincipal, err := authHandler.userRepository.FindOrCreateUser(profile)
	if err != nil {
		logIfErr(err)
		ErrorResponse(w, 500, "something went wrong")
		return
	}

	token, err := authHandler.jwtutil.SignToken(userprincipal)
	if err != nil {
		logIfErr(err)
		ErrorResponse(w, 500, "something went wrong")
		return
	}

	response := models.TokenResponse{JwtToken: token}

	encodeToJson(w, 200, response)
}

func (authHandler AuthHandler) ValidateTokenTillNextDay(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Token string `json:"jwtToken"`
	}{}

	json.NewDecoder(r.Body).Decode(&body)
	token := body.Token

	isValidTillNextDay := authHandler.jwtutil.ValidTillNextDay(token)
	if !isValidTillNextDay {
		ErrorResponse(w, 400, "Not a valid Token")
		return
	}
	resp := struct{}{}
	encodeToJson(w, 200, resp)
}

// Utilities
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
		return models.GoogleProfileInfo{}, errors.New("could not get profile data")
	}

	googleProfileInfo := models.GoogleProfileInfo{}
	err = json.NewDecoder(res.Body).Decode(&googleProfileInfo)
	if err != nil {
		return models.GoogleProfileInfo{}, err
	}

	return googleProfileInfo, nil
}

func getAccessTokenFromCode(client *http.Client, tokenReq models.AccessTokenReqGoogle) (models.AccessTokenRespGoogle, error) {

	log.Println("getting access token")
	url := "https://oauth2.googleapis.com/token"
	method := http.MethodPost

	log.Println(tokenReq.AuthCode)

	log.Println(tokenReq)

	payload, err := json.Marshal(tokenReq)
	// log.Println(payload)

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
	log.Println(res.StatusCode)
	buf := new(strings.Builder)
	body := buf.String()
	log.Println(body)
	log.Print(res)
	log.Println(res.StatusCode != http.StatusOK)

	if res.StatusCode != http.StatusOK {
		return models.AccessTokenRespGoogle{}, errors.New("get token failed")
	}

	accessTokenReponse := models.AccessTokenRespGoogle{}
	err = json.NewDecoder(res.Body).Decode(&accessTokenReponse)
	if err != nil {
		return models.AccessTokenRespGoogle{}, err
	}
	return accessTokenReponse, nil

}

func logIfErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
