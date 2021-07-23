package repository

import (
	"github.com/mananwalia959/go-todos-app/pkg/models"
)

type UserRepository interface {
	FindOrCreateUser(models.GoogleProfileInfo) (models.UserPrincipal, error)
}
