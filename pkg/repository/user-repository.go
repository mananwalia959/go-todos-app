package repository

import (
	"github.com/google/uuid"
	"github.com/mananwalia959/go-todos-app/pkg/models"
)

type UserRepository interface {
	FindOrCreateUser(models.GoogleProfileInfo) (models.UserPrincipal, error)
}

type UserRepositoryImpl struct {
	userStore map[string]uuid.UUID
}

func InitializeUserRepository() UserRepository {
	return &UserRepositoryImpl{userStore: map[string]uuid.UUID{}}
}

func (repo *UserRepositoryImpl) FindOrCreateUser(googleProfile models.GoogleProfileInfo) (models.UserPrincipal, error) {
	userId, present := repo.userStore[googleProfile.Email]
	if present {
		return models.UserPrincipal{Id: userId,
			Email:   googleProfile.Email,
			Name:    googleProfile.Name,
			Picture: googleProfile.Picture,
		}, nil
	}

	newUserId := uuid.New()
	repo.userStore[googleProfile.Email] = newUserId

	return models.UserPrincipal{Id: newUserId,
		Email:   googleProfile.Email,
		Name:    googleProfile.Name,
		Picture: googleProfile.Picture,
	}, nil
}
