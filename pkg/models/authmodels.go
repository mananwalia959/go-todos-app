package models

type TokenRequest struct {
	Code string `json:"code"`
}

type AccessTokenRespGoogle struct {
	AccessToken string `json:"access_token"`
}

type AccessTokenReqGoogle struct {
	Client_id     string `json:"client_id"`
	Client_secret string `json:"client_secret"`
	Grant_type    string `json:"grant_type"`
	Redirect_uri  string `json:"redirect_uri"`
	AuthCode      string `json:"code"`
}

type Profile struct {
	Id      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}
