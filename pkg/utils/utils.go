package utils

import (
	"context"
	"net/http"

	"github.com/mananwalia959/go-todos-app/pkg/models"
)

func GetUserPrincipal(r *http.Request) models.UserPrincipal {
	return GetUserPrincipalFromContext(r.Context())
}

func GetUserPrincipalFromContext(ctx context.Context) models.UserPrincipal {
	var key models.UserPrincipalCtxKey = "UserPrincipal"
	upi := ctx.Value(key)
	if upi == nil {
		panic("empty userPrincipal ")
	}
	up, err := upi.(models.UserPrincipal)
	if !err {
		panic("can't cast to userPrincipal")
	}
	return up

}
