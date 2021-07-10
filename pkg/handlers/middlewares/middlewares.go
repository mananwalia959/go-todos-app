package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mananwalia959/go-todos-app/pkg/handlers"
	"github.com/mananwalia959/go-todos-app/pkg/models"
	"github.com/mananwalia959/go-todos-app/pkg/oauth"
)

func GetAuthMiddleWare(jwtUtil oauth.JWTUtil) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			up, err := createUserPrincipalFromToken(r, jwtUtil)
			if err != nil {
				handlers.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized , provide a valid jwt token")
				return
			}

			var key models.UserPrincipalCtxKey = "UserPrincipal"
			ctx := context.WithValue(r.Context(), key, up)
			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}

}

func createUserPrincipalFromToken(r *http.Request, jwtUtil oauth.JWTUtil) (models.UserPrincipal, error) {
	authorizationVal := r.Header.Get("Authorization")

	if !strings.HasPrefix(authorizationVal, "Bearer ") {
		return models.UserPrincipal{}, errors.New("no valid token")
	}

	token := strings.TrimPrefix(authorizationVal, "Bearer ")
	return jwtUtil.VerifySign(token)
}

func GetUserPrincipal(r *http.Request) models.UserPrincipal {
	var key models.UserPrincipalCtxKey = "UserPrincipal"
	upi := r.Context().Value(key)
	if upi == nil {
		panic("empty userPrincipal ")
	}
	up, err := upi.(models.UserPrincipal)
	if !err {
		panic("can't cast to userPrincipal")
	}
	return up
}

func PanicRecovermiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				handlers.ErrorResponse(w, 500, "Something went wrong")
			}
		}()

		next.ServeHTTP(w, r)

	})

}

// (w http.ResponseWriter, r *http.Request) {
