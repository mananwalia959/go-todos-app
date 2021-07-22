package config

type Appconfig struct {
	OauthClientId     string
	OauthClientSecret string
	OauthRedirectUrl  string
	SecretKeyJwt      string
	PostgresUrl       string
	PostgresDbName    string
	PostgresUsername  string
	PostgresPassword  string
}
