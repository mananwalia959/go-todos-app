package oauth

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mananwalia959/go-todos-app/pkg/config"
	"github.com/mananwalia959/go-todos-app/pkg/models"
	"github.com/pascaldekloe/jwt"
)

type JWTUtil interface {
	VerifySign(string) (models.UserPrincipal, error)
	SignToken(models.UserPrincipal) (string, error)
	ValidTillNextDay(string) bool
}

func InitializeJwtUtil(appconfig *config.Appconfig) JWTUtil {
	return &JWTUtilImpl{secret: []byte(appconfig.SecretKeyJwt)}
}

type JWTUtilImpl struct {
	secret []byte
}

func (jwtUtilImpl *JWTUtilImpl) VerifySign(token string) (models.UserPrincipal, error) {
	return tokenValidate(token, time.Now(), jwtUtilImpl.secret)

}

func (jwtUtilImpl *JWTUtilImpl) SignToken(userprincipal models.UserPrincipal) (string, error) {
	claims := jwt.Claims{}

	claims.Subject = userprincipal.Id.String()
	claims.Set = map[string]interface{}{"email": userprincipal.Email,
		"picture": userprincipal.Picture,
		"name":    userprincipal.Name}

	claims.Issued = jwt.NewNumericTime(time.Now().Round(time.Second))
	claims.Expires = jwt.NewNumericTime(time.Now().AddDate(0, 0, 10).Round(time.Second))
	tokenbytes, err := claims.HMACSign(jwt.HS512, jwtUtilImpl.secret)

	if err != nil {
		return "", err
	}
	return string(tokenbytes), nil

}

func (jwtUtilImpl *JWTUtilImpl) ValidTillNextDay(token string) bool {
	_, err := tokenValidate(token, time.Now().AddDate(0, 0, 1), jwtUtilImpl.secret)
	return err == nil
}

func tokenValidate(token string, time time.Time, secret []byte) (models.UserPrincipal, error) {

	claims, err := jwt.HMACCheck([]byte(token), secret)
	if err != nil {
		return models.UserPrincipal{}, err
	}

	if claims.Valid(time) {
		return parseClaimsForUserPrincipal(claims)

	} else {
		return models.UserPrincipal{}, errors.New("token expired")
	}
}

func parseClaimsForUserPrincipal(claims *jwt.Claims) (models.UserPrincipal, error) {
	userprincipal := models.UserPrincipal{}
	userprincipal.Id = uuid.MustParse(claims.Subject)
	userprincipal.Name = castInterfaceToString(claims.Set["name"])
	userprincipal.Email = castInterfaceToString(claims.Set["email"])
	userprincipal.Picture = castInterfaceToString(claims.Set["picture"])

	return userprincipal, nil
}
func castInterfaceToString(s interface{}) string {
	return fmt.Sprintf("%v", s)
}
