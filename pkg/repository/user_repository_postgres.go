package repository

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mananwalia959/go-todos-app/pkg/models"
)

func InitializePostgresUserRepository(pool *pgxpool.Pool) UserRepository {
	return &UserRepositoryPostgresImpl{pool: pool}
}

type UserRepositoryPostgresImpl struct {
	pool *pgxpool.Pool
}

func (repo *UserRepositoryPostgresImpl) FindOrCreateUser(googleProfile models.GoogleProfileInfo) (models.UserPrincipal, error) {

	rows, err := repo.pool.Query(context.Background(), "select id from users where email = $1", googleProfile.Email)
	if err != nil {
		log.Println(err)
		return models.UserPrincipal{}, err
	}

	defer rows.Close()
	if rows.Next() {
		userid := uuid.UUID{}
		err = rows.Scan(&userid)
		if err != nil { // yuck
			return models.UserPrincipal{}, err
		}
		return getUserPrincipal(googleProfile, userid), nil
	}
	userid := uuid.UUID{}
	err = repo.pool.QueryRow(context.Background(), "insert into users (email ) values ($1) returning id", googleProfile.Email).Scan(&userid)
	if err != nil {
		return models.UserPrincipal{}, err
	}
	return getUserPrincipal(googleProfile, userid), nil

}

func getUserPrincipal(googleProfile models.GoogleProfileInfo, userid uuid.UUID) models.UserPrincipal {
	return models.UserPrincipal{Id: userid,
		Email:   googleProfile.Email,
		Name:    googleProfile.Name,
		Picture: googleProfile.Picture,
	}
}
